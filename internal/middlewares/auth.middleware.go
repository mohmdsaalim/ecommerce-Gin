package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"No token avilable"})
			return 
		}

		// 2. Expect format: Bearer <token>
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			c.Abort()
			return
		}
		tokenstring := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, jwt.ErrSignatureInvalid
			} 
			return []byte(config.AppConfig.JWT.Secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid token claims"})
			c.Abort()
			return 
		}
		
		// Extract claims safely
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		// Get user_id
		userID, ok := claims["user_id"].(float64) // JWT stores numbers as float64
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user_id missing in token",
			})
			c.Abort()
			return
		}

		// Get role
		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "role missing in token",
			})
			c.Abort()
			return
		}

		//  Attach to context
		c.Set("userID", uint(userID))
		c.Set("role", role)
		c.Next()
		
	}
}