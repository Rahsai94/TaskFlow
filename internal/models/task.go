package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"` // "To-Do", "In Progress", "Completed"
    AssignedTo  uint   `json:"assigned_to"` // User ID
}
