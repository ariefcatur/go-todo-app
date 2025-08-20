package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/controllers"
	"go-todo-app/middlewares"
	"go-todo-app/models"
	"log"
	"net/http"
	"time"
)

func main() {
	//config.InitDatabase()
	config.Load()
	config.ConnectDB()

	// Auto Migrate
	err := config.DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Error migrating DB")
	}

	router := gin.New()
	router.Use(gin.Recovery())
	for _, m := range middlewares.Security() {
		router.Use(m)
	}

	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	//protected := r.Group("/api")
	//protected.Use(middlewares.AuthMiddleware())
	//{
	//	protected.POST("/tasks", controllers.CreateTask)
	//	protected.GET("/tasks", controllers.GetTasks)
	//	protected.PUT("/tasks/:id", controllers.UpdateTask)
	//	protected.DELETE("/tasks/:id", controllers.DeleteTask)
	//}

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
	log.Fatal(srv.ListenAndServe())
}
