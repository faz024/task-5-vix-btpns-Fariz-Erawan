package middlewares

import (
	"strings"

	"task-5-vix-btpns-fariz/app/auth"

	"github.com/gin-gonic/gin"
)

// function to protect routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization") //Get bearer token
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		err := auth.ValidateToken(strings.Split(tokenString, "Bearer ")[1]) //Validate token
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
