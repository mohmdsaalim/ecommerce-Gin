// Package database is responsible for initializing and providing
// the PostgreSQL database connection.
package database

import (
	"fmt"
	"log"

	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a shared GORM database instance used across the application
// after a successful connection is established.
var DB *gorm.DB

// ConnectPostgres initializes a PostgreSQL connection using
// configuration values and assigns it to the global DB variable.
func ConnectPostgres() {
	// Load application configuration
	cfg := config.AppConfig

	// Construct PostgreSQL DSN (Data Source Name)
	// This contains all required credentials and connection details
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.Port,
	)

	// Open connection using GORM and PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { // Stop application if database connection fails
		log.Fatalf("[ERROR] failed to connect to PostgreSQL: %v", err)
	}

	// Automigrate 
	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductVariant{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},	
		&models.OrderItem{},
		&models.Address{},

	)
	if err != nil {
		log.Fatal("Auto-migrate Failed:", err)//Automigrate handlng
	}
	// Store the connection for reuse throughout the app
	DB = db

	// Log successful database connection
	log.Println("✅ Connected to PostgreSQL")
}
