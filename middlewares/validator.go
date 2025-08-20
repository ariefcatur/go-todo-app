package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func BindAndValidate[T any](c *gin.Context, dst *T) bool {
	if err := c.ShouldBindJSON(dst); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return false
	}
	if err := validate.Struct(dst); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return false
	}
	return true
}
