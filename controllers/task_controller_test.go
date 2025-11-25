package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/controllers"
	"go-todo-app/helpers"
	"go-todo-app/internal/testutil"
	"go-todo-app/models"
)

func setupTaskRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	config.C.JWTSecret = "testsecret"
	config.C.JWTExpiry = time.Hour
	config.DB = testutil.NewTestDB()

	// seed one user
	u := models.User{Username: "u1", Email: "u1@example.com", PasswordHash: "x"}
	if err := config.DB.Create(&u).Error; err != nil {
		panic(err)
	}
	token, _ := helpers.CreateAccessToken(u.ID)

	// auth middleware mock (sederhana): set user_id langsung
	auth := func(c *gin.Context) {
		c.Set("user_id", u.ID)
		c.Next()
	}

	api := r.Group("/api")
	api.Use(auth)
	api.POST("/tasks", controllers.CreateTask)
	api.GET("/tasks", controllers.GetTasks)
	api.PUT("/tasks/:id", controllers.UpdateTask)
	api.DELETE("/tasks/:id", controllers.DeleteTask)

	// simpan token di context test (hack: header di request)
	r.Use(func(c *gin.Context) {
		c.Request.Header.Set("Authorization", "Bearer "+token)
		c.Next()
	})

	return r
}

func TestTaskCRUD(t *testing.T) {
	r := setupTaskRouter()

	// create
	body, _ := json.Marshal(map[string]string{
		"title":    "Belajar Go",
		"priority": "high",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", w.Code, w.Body.String())
	}

	// list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/tasks?priority=high", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("list status=%d body=%s", w.Code, w.Body.String())
	}

	// verify pagination structure
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if _, ok := resp["data"].(map[string]interface{})["pagination"]; !ok {
		t.Fatal("pagination field missing in response")
	}
}
