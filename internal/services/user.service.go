package services

import (
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type UserService struct {
	repo repositories.Repository
}

func NewUserService(repo repositories.Repository) *UserService {
	return &UserService{repo: repo}
}

// get the profile
func (s *UserService) GetProfile(userID uint) (*models.User, error) {
	var user models.User

	err := s.repo.FindOne(
		&user,
		"id = ?",
		[]string{"Addresses"},
		userID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// updtae the profile
func (s *UserService) UpdateProfile(userID uint, data map[string]interface{}) error {
	return s.repo.UpdateFields(&models.User{}, userID, data)
}

// Address service
func (s *UserService) AddAddress(userID uint, address *models.Address) error {
	address.UserID = userID
	return s.repo.Insert(address)
}