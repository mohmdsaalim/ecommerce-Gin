package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/routes"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
)

func main() {
	// Load YAML config
	config.LoadConfig()

	// Connect to PostgreSQL & Redis
	database.ConnectPostgres()
	// database.ConnectRedis()

	// pedning reddis connection -----

	// Initialize  Genric repo 
	repo := repositories.NewPgSQLRepository()

	// Initialize Services
	authService := services.NewAuthService(repo)

	//  Initialize Gin Engine
	r := gin.Default()

	// call register router  
	routes.RegisterRoute(r, authService)
	
	//  Start server
	port := config.AppConfig.App.Port
	log.Printf("🚀 Server running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}