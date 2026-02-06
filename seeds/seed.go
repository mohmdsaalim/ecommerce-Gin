package seeds

import (
	"log"

	"gorm.io/gorm"
)

// RunSeeds runs all seeders in proper order
func RunSeeds(db *gorm.DB) error {
	log.Println("Running database seeds...")

	if err := SeedUsers(db); err != nil{
		return err
	}// completed 
	
	 if err := SeedProducts(db); err != nil{
		return err
	 }
	return nil
}