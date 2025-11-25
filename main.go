package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/controllers"
	"go-todo-app/middlewares"
	"go-todo-app/models"
)

func main() {
	config.Load()
	gin.SetMode(config.C.GinMode)
	config.ConnectDB()

	// Auto Migrate
	err := config.DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Error migrating DB")
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.RequestID())
	router.Use(middlewares.StructuredLogger())
	for _, m := range middlewares.Security() {
		router.Use(m)
	}

	// Health check
	router.GET("/health", controllers.HealthCheck)

	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	api := router.Group("/api")
	api.Use(middlewares.JWTAuth())
	api.GET("/tasks", controllers.GetTasks)
	api.POST("/tasks", controllers.CreateTask)
	api.PUT("/tasks/:id", controllers.UpdateTask)
	api.DELETE("/tasks/:id", controllers.DeleteTask)

	srv := &http.Server{
		Addr:         ":" + config.C.Port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Server starting on port %s", config.C.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
