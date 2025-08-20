package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-todo-app/config"
	"go-todo-app/helpers"
	"net/http"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
			return
		}

		claims := &helpers.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenInvalidIssuer
			}
			return []byte(config.C.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
