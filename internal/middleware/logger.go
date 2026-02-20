package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context){
		start := time.Now()
		c.Next() //c.Next() is a crucial part of the middleware function. It tells Gin to proceed with the next handler in the chain, which could be another middleware or the final route handler. This allows the middleware to perform actions both before and after the request is processed, such as logging the start time before and calculating the duration after.
		duration := time.Since(start)
		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
	)
	}
}