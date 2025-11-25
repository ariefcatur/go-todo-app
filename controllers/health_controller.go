package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-todo-app/config"
)

func HealthCheck(c *gin.Context) {
	// Check database connection
	sqlDB, err := config.DB.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "database connection error",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "database ping failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"service":   "go-todo-app",
	})
}
