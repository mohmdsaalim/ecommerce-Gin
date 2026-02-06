package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/app"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
)

func main() {
	config.LoadConfig()
	database.ConnectPostgres()
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	app.RegisterDependencies(r)

	port := config.AppConfig.App.Port
	log.Printf("🚀 Server running at http://localhost:%s", port)
	r.Run(":" + port)
}