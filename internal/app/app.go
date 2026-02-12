package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/middlewares"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/routes"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
)

func RegisterDependencies(r *gin.Engine) {
	r.Use(cors.New(middlewares.CORSMiddleware()))
	// repo
	repo := repositories.NewPgSQLRepository()

	// services
	authService := services.NewAuthService(repo, database.RedisClient)
	productService := services.NewProductService(repo)
	cartService := services.NewCartService(repo)
	orderService := services.NewOrderService(repo)
	userSrevice := services.NewUserService(repo)
	adminService := services.NewAdminService(repo)
	adminProductService := services.NewAdminProductService(repo)
	adminUserService := services.NewAdminUserService(repo)
	adminOrderService := services.NewAdminOrderService(repo)

	// routes
	routes.RegisterRoute(
		r, // passed gin Engine
		authService, // auth service
		productService, // productService
		cartService,// cart service
		orderService, // orderservice
		userSrevice,
		adminService,
		adminProductService,
		adminUserService,
		adminOrderService,
		)
}