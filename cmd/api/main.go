package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/app"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services/workers"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
)

func main() {

	config.LoadConfig()// config loading

	database.ConnectRedis()// connect reddis
	database.ConnectPostgres() // db connection 
	r := gin.Default() // Gin Engine

	go workers.StartOTPWorker()
	app.RegisterDependencies(r)

	port := config.AppConfig.App.Port
	log.Printf("🚀 Server running at http://localhost:%s", port)
	r.Run(":" + port)

}


// pending ....
// redis conection ✅

// email template
// sendOTP
// verifyOTP
// forgott password
// rate limiting
// workers  

// shutdown 
// prs core
// imbel 