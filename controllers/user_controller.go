package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-todo-app/config"
	"go-todo-app/helpers"
	"go-todo-app/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation Error", gin.H{
			"details": err.Error(),
		})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server Error", gin.H{"details": "Failed to hash password"})
		return
	}

	// Buat user baru
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server Error", gin.H{"details": "Failed to create user"})
		return
	}

	helpers.APIResponse(c, http.StatusCreated, "User created successfully", gin.H{
		"user_id":  user.ID,
		"username": user.Username,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation Error", gin.H{
			"details": err.Error(),
		})
		return
	}

	var user models.User

	// Cari user berdasarkan username
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", gin.H{
			"details": "Invalid credentials",
		})
		return
	}

	// Verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", gin.H{
			"details": "Invalid credentials",
		})
		return
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{
			"details": "Failed to generate token",
		})
		return
	}

	helpers.APIResponse(c, http.StatusOK, "Login successful", gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}
