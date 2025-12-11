package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/logger"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		switch {
		case status >= 500:
			logger.Error("%s %s | %d | %v", method, path, status, latency)
		case status >= 400:
			logger.Warn("%s %s | %d | %v", method, path, status, latency)
		default:
			logger.Info("%s %s | %d | %v", method, path, status, latency)
		}
	}
}
