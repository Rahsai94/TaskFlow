package handlers

import (
    "github.com/gin-gonic/gin"
    "taskflow/config"
    "taskflow/models"
    "taskflow/utils"
    "net/http"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password and save user
    user.Password = utils.HashPassword(user.Password)
    config.DB.Create(&user)
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    config.DB.Where("username = ?", input.Username).First(&user)

    if user.ID == 0 || !utils.CheckPassword(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := utils.GenerateToken(user.ID)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
