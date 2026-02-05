package seeds

import (
	"errors"
	"log"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	adminPass, err := utils.HashPassword("admin@123")
	if err != nil {
		return err
	}

	userPass, err := utils.HashPassword("user@123")
	if err != nil {
		return err
	}

	users := []models.User{
		{
			Email:         "admin@example.com",
			Username:      "admin",
			PasswordHash:  adminPass,
			Role:          "admin",
			Status:        "active",
			EmailVerified: true,
		},
		{
			Email:         "user@example.com",
			Username:      "user",
			PasswordHash:  userPass,
			Role:          "user",
			Status:        "active",
			EmailVerified: false,
		},
	}

	for _, user := range users {
		var existing models.User

		err := db.Where("email = ?", user.Email).First(&existing).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {

			if err := db.Create(&user).Error; err != nil {
				return err
			}

			log.Println("Seeded user:", user.Email)

		} else if err != nil {
			// Handles unexpected DB errors
			return err
		}
	}

	return nil
}