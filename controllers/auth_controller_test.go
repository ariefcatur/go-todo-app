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
	"go-todo-app/internal/testutil"
)

func setupAuthRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// set config for test
	config.C.JWTSecret = "testsecret"
	config.C.JWTExpiry = 30 * time.Minute
	config.DB = testutil.NewTestDB()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	return r
}

func TestRegisterAndLogin(t *testing.T) {
	r := setupAuthRouter()

	// register
	regBody, _ := json.Marshal(map[string]string{
		"username": "tester",
		"email":    "tester@example.com",
		"password": "pass12345",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(regBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("register status = %d, body=%s", w.Code, w.Body.String())
	}

	// login
	loginBody, _ := json.Marshal(map[string]string{
		"identity": "tester@example.com",
		"password": "pass12345",
	})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(loginBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("login status = %d, body=%s", w.Code, w.Body.String())
	}
}
