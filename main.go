package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-todo-app/config"
	"go-todo-app/controllers"
	"go-todo-app/middlewares"
	"go-todo-app/models"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDatabase()

	// Auto Migrate
	config.DB.AutoMigrate(&models.User{}, &models.Task{})

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
