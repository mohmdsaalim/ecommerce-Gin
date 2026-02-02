package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/controllers"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)


func RegisterRoute(r *gin.Engine , authService *services.AuthService) {
	authController := controllers.NewAuthController(authService)

// checking route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status":"OK"})
	})
// PUBLIC routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}
}