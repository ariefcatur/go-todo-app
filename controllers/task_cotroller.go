package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/helpers"
	"go-todo-app/models"
	"net/http"
)

func CreateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{
			"details": err.Error(),
		})
		return
	}

	task.UserID = userID.(uint)
	if err := config.DB.Create(&task).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{
			"details": "Failed to create task",
		})
		return
	}

	helpers.APIResponse(c, http.StatusCreated, "Task created successfully", task)
}

func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var tasks []models.Task

	status := c.Query("status")
	priority := c.Query("priority")

	query := config.DB.Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	if err := query.Find(&tasks).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{
			"details": "Failed to fetch tasks",
		})
		return
	}

	helpers.APIResponse(c, http.StatusOK, "Tasks retrieved successfully", tasks)
}

func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID := c.Param("id")

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Task not found", gin.H{
			"details": err.Error(),
		})
		return
	}

	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{
			"details": err.Error(),
		})
		return
	}

	config.DB.Model(&task).Updates(updateTask)
	helpers.APIResponse(c, http.StatusOK, "Task updated successfully", task)
}

func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID := c.Param("id")

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Task not found", gin.H{
			"details": err.Error(),
		})
		return
	}

	config.DB.Delete(&task)
	helpers.APIResponse(c, http.StatusOK, "Task deleted successfully", nil)
}
