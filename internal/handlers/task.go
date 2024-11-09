package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "taskflow/config"
    "taskflow/models"
)

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    config.DB.Find(&tasks)
    c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&task)
    c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := config.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&task)
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.Task{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
