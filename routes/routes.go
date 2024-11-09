package routes

import (
    "github.com/gin-gonic/gin"
    "taskflow/handlers"
    "taskflow/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    authorized := r.Group("/")
    authorized.Use(middlewares.AuthMiddleware())
    authorized.GET("/tasks", handlers.GetTasks)
    authorized.POST("/tasks", handlers.CreateTask)
    authorized.PUT("/tasks/:id", handlers.UpdateTask)
    authorized.DELETE("/tasks/:id", handlers.DeleteTask)
	authorized.GET("/ws/tasks", handlers.TaskWebSocket)


    return r
}
