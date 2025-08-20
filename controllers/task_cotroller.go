package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/helpers"
	"go-todo-app/models"
)

func CreateTask(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var in struct {
		Title       string `json:"title" binding:"required,min=1"`
		Description string `json:"description"`
		Priority    string `json:"priority"` // low|medium|high
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": err.Error()})
		return
	}
	p := strings.ToLower(strings.TrimSpace(in.Priority))
	if p == "" {
		p = "medium"
	}
	if p != "low" && p != "medium" && p != "high" {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": "priority must be low|medium|high"})
		return
	}
	task := models.Task{
		UserID:      uid.(int64),
		Title:       strings.TrimSpace(in.Title),
		Description: strings.TrimSpace(in.Description),
		Priority:    p,
	}
	if err := config.DB.Create(&task).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{"details": "fail create task"})
		return
	}
	helpers.APIResponse(c, http.StatusCreated, "Task created", task)
}

func GetTasks(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var tasks []models.Task
	q := config.DB.Where("user_id = ?", uid.(int64))

	if s := c.Query("status"); s != "" {
		s = strings.ToLower(s)
		if s != "pending" && s != "completed" {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": "status must be pending|completed"})
			return
		}
		q = q.Where("status = ?", s)
	}
	if p := c.Query("priority"); p != "" {
		p = strings.ToLower(p)
		if p != "low" && p != "medium" && p != "high" {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": "priority must be low|medium|high"})
			return
		}
		q = q.Where("priority = ?", p)
	}
	if err := q.Order("id desc").Find(&tasks).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{"details": "fail query"})
		return
	}
	helpers.APIResponse(c, http.StatusOK, "OK", tasks)
}

func UpdateTask(c *gin.Context) {
	uid, _ := c.Get("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", id, uid.(int64)).First(&task).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Not found", gin.H{"details": "task not found"})
		return
	}
	var in struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *string `json:"status"`   // pending|completed
		Priority    *string `json:"priority"` // low|medium|high
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": err.Error()})
		return
	}
	if in.Title != nil {
		task.Title = strings.TrimSpace(*in.Title)
	}
	if in.Description != nil {
		task.Description = strings.TrimSpace(*in.Description)
	}
	if in.Status != nil {
		s := strings.ToLower(strings.TrimSpace(*in.Status))
		if s != "pending" && s != "completed" {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": "status must be pending|completed"})
			return
		}
		task.Status = s
	}
	if in.Priority != nil {
		p := strings.ToLower(strings.TrimSpace(*in.Priority))
		if p != "low" && p != "medium" && p != "high" {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{"details": "priority must be low|medium|high"})
			return
		}
		task.Priority = p
	}
	if err := config.DB.Save(&task).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{"details": "fail update"})
		return
	}
	helpers.APIResponse(c, http.StatusOK, "Updated", task)
}

func DeleteTask(c *gin.Context) {
	uid, _ := c.Get("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := config.DB.Where("id = ? AND user_id = ?", id, uid.(int64)).Delete(&models.Task{}).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{"details": "fail delete"})
		return
	}
	helpers.APIResponse(c, http.StatusOK, "Deleted", gin.H{"id": id})
}
