package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/middlewares"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/routes"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

func RegisterDependencies(r *gin.Engine) {
	r.Use(cors.New(middlewares.CORSMiddleware()))
	// repo
	repo := repositories.NewPgSQLRepository()

	// services
	authService := services.NewAuthService(repo)
	productService := services.NewProductService(repo)
	cartService := services.NewCartService(repo)
	orderService := services.NewOrderService(repo)
	userSrevice := services.NewUserService(repo)
	
	// routes
	routes.RegisterRoute(
		r, // passed gin Engine
		authService, // auth service
		productService, // productService
		cartService,// cart service
		orderService, // orderservice
		userSrevice,
		)
}