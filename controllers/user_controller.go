package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/config"
	"go-todo-app/helpers"
	"go-todo-app/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required,min=3,max=30"`
		Email    string `json:"email"    binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation Error", gin.H{
			"details": err.Error(),
		})
		return
	}

	input.Username = strings.TrimSpace(input.Username)
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))

	// Validate Email
	if !helpers.IsValidEmail(input.Email) {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{
			"details": "Invalid email format",
		})
		return
	}

	// Check if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{
			"details": "Email already registered",
		})
	}

	// Check if username already exists
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation error", gin.H{
			"details": "Username already taken",
		})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server Error", gin.H{"details": "Failed to hash password"})
		return
	}

	// Create New User
	user := models.User{
		Username:     input.Username,
		Email:        strings.ToLower(input.Email),
		PasswordHash: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server Error", gin.H{"details": "Failed to create user"})
		return
	}

	helpers.APIResponse(c, http.StatusCreated, "User created successfully", gin.H{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Identity string `json:"identity" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Validation Error", gin.H{
			"details": err.Error(),
		})
		return
	}

	identity := strings.TrimSpace(input.Identity)
	var user models.User

	// Check if identity is email or username
	q := config.DB
	if helpers.IsValidEmail(identity) {
		q = q.Where("email = ?", strings.ToLower(identity))
	} else {
		q = q.Where("username = ?", identity)
	}

	if err := q.First(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", gin.H{
			"details": "Invalid credentials",
		})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", gin.H{
			"details": "Invalid credentials",
		})
		return
	}

	token, err := helpers.CreateAccessToken(user.ID) // <- konsisten jwt/v5
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Server error", gin.H{"details": "Failed to generate token"})
		return
	}

	helpers.APIResponse(c, http.StatusOK, "Login successful", gin.H{
		"token_type": "Bearer",
		"expires_in": int(config.C.JWTExpiry.Seconds()),
		"token":      token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
