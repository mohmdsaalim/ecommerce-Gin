package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequareRole(requiredRole string) gin.HandlerFunc {
	return  func(c *gin.Context) {
		// authmiddle ware run before this
		role := c.GetString("role")

		if role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":"role not found in context",
			})
			c.Abort()
			return 
		}
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error":"acces denied",
			})
			c.Abort()
			return 
		}
		c.Next()
	}
}