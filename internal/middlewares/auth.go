package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "taskflow/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" || !utils.ValidateToken(token) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
