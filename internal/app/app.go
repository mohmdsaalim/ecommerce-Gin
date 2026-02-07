package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/routes"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

func RegisterDependencies(r *gin.Engine) {
	// repo
	repo := repositories.NewPgSQLRepository()

	// services
	authService := services.NewAuthService(repo)
	productService := services.NewProductService(repo)
	cartService := services.NewCartService(repo)
	// routes
	routes.RegisterRoute(
		r, // passed gin Engine
		authService, // auth service
		productService, // productService
		cartService,
		)
}