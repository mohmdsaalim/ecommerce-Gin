package services

import (
	"errors"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type AdminUserService struct {
	repo repositories.Repository
}

func NewAdminUserService(repo repositories.Repository) *AdminUserService {
	return &AdminUserService{repo: repo}
}

func (s *AdminUserService) GetAllUsers() ([]models.User, error) {

	var users []models.User

	err := s.repo.FindAll(
		&users,
		"",
		"created_at DESC",
		[]string{},
	)

	return users, err
}

func (s *AdminUserService) GetUserByID(id uint) (*models.User, error) {

	var user models.User

	err := s.repo.FindByID(&user, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AdminUserService) DeactivateUser(id uint) error {

	return s.repo.UpdateFields(
		&models.User{},
		id,
		map[string]interface{}{
			"status": "inactive",
		},
	)
}


func (s *AdminUserService) ActivateUser(id uint) error {

	return s.repo.UpdateFields(
		&models.User{},
		id,
		map[string]interface{}{
			"status": "active",
		},
	)
}

func (s *AdminUserService) ChangeRole(id uint, role string) error {

	if role != "admin" && role != "user" {
		return errors.New("invalid role")
	}

	return s.repo.UpdateFields(
		&models.User{},
		id,
		map[string]interface{}{
			"role": role,
		},
	)
}

func (s *AdminUserService) DeleteUser(id uint) error {

	return s.repo.Delete(
		&models.User{},
		"id = ?",
		id,
	)
}