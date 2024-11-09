package main

import (
    "taskflow/config"
    "taskflow/routes"
)

func main() {
    // Initialize configuration
    db := config.ConnectDatabase()
	db.AutoMigrate(&models.User{}, &models.Task{})
    defer db.Close()

    // Setup routes
    r := routes.SetupRouter()

    // Start server
    r.Run(":8080")
}

