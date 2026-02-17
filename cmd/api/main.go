
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/app"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services/workers"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
)

func main() {

	// Load Config
	config.LoadConfig()

	// Connect Databases
	database.ConnectRedis()
	database.ConnectPostgres()

	// Create Gin Engine
	r := gin.Default()

	// Register routes & dependencies
	app.RegisterDependencies(r)

	// Start background worker
	go workers.StartOTPWorker()

	port := config.AppConfig.App.Port

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server in goroutine
	go func() {
		log.Printf("🚀 Server running at http://localhost:%s", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited properly")
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


 // =---------------- with out graceful shodown-----------------------
// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/mohmdsaalim/ecommerce-Gin/config"
// 	"github.com/mohmdsaalim/ecommerce-Gin/internal/app"
// 	"github.com/mohmdsaalim/ecommerce-Gin/internal/services/workers"
// 	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
// )

// func main() {

// 	config.LoadConfig()// config loading

// 	database.ConnectRedis()// connect reddis
// 	database.ConnectPostgres() // db connection 
// 	r := gin.Default() // Gin Engine

// 	go workers.StartOTPWorker()
// 	app.RegisterDependencies(r)

// 	port := config.AppConfig.App.Port
// 	log.Printf("🚀 Server running at http://localhost:%s", port)
// 	r.Run(":" + port)

// }
