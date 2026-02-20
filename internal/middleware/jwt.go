package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")
		if token != "mysecrettoken" {
			c.AbortWithStatusJSON(401,gin.H{
				"error":"unauthorized",
			})
			return
		}
	c.Next()
}
}