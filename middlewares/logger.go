package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func StructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		duration := time.Since(start)
		requestID := c.GetString("request_id")

		log.Printf("[%s] %s %s | status=%d | duration=%v | request_id=%s | ip=%s",
			c.Request.Method,
			path,
			query,
			c.Writer.Status(),
			duration,
			requestID,
			c.ClientIP(),
		)
	}
}
