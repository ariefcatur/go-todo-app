package middlewares

import (
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func Security() []gin.HandlerFunc {
	// Security headers
	secHeaders := func(c *gin.Context) {
		h := c.Writer.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("X-XSS-Protection", "0")
		h.Set("Referrer-Policy", "no-referrer")
		h.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none';")
		c.Next()
	}

	// Rate limit: 100 req/jam per IP (contoh)
	//limiter := tollbooth.NewLimiter(100, &tollbooth.{
	//	ExpirationTTL: time.Hour,
	//})
	limiter := tollbooth.NewLimiter(100, nil)

	return []gin.HandlerFunc{
		secHeaders,
		tollbooth_gin.LimitHandler(limiter),
	}
}
