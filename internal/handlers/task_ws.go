package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

var wsUpgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func TaskWebSocket(c *gin.Context) {
    conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to establish WebSocket connection"})
        return
    }
    defer conn.Close()

    clients[conn] = true

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            delete(clients, conn)
            break
        }

        if messageType == websocket.TextMessage {
            broadcast <- string(message)
        }
    }
}

func BroadcastTaskUpdates() {
    for {
        message := <-broadcast
        for client := range clients {
            if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
                client.Close()
                delete(clients, client)
            }
        }
    }
}
