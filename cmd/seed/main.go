package main

import (
	"log"

	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
	"github.com/mohmdsaalim/ecommerce-Gin/seeds"
)

// var DB *gorm.DB

func main()  {
	config.LoadConfig()
	database.ConnectPostgres()

	if err := seeds.RunSeeds(database.DB); err != nil{
		log.Fatalf("seeding failed: %v", err)
	}
	log.Println("Database seeded successfully")
}