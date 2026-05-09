package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[%s] %s %s %d %s",
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			status,
			duration,
		)
	}
}
