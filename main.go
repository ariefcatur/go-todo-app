package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/controllers"
	"go-todo-app/middlewares"
	"go-todo-app/models"
	"log"
)

func main() {
	config.InitDatabase()

	// Auto Migrate
	err := config.DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Error migrating DB")
	}

	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/tasks", controllers.CreateTask)
		protected.GET("/tasks", controllers.GetTasks)
		protected.PUT("/tasks/:id", controllers.UpdateTask)
		protected.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	r.Run(":8080")
}
