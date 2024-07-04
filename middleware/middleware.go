// Package middleware provides the middleware for the application.
package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a middleware that sets the headers for CORS requests.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Set the Access-Control-Allow-Origin header to the appropriate value.
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, userID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// AuthMiddleware is a middleware that checks for the userID header.
// Note: This is a placeholder for a real authentication middleware.
// TODO: Implement a real authentication middleware.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("userID")

		if userID == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
