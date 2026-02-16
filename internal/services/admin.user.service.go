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

// GetAllUsers retrieves all users with pagination for the admin panel
func (s *AdminUserService) GetAllUsers(page, limit int) ([]models.User, error) {

	var users []models.User

	// Calculate offset for pagination
	// page 1 -> offset 0
	// page 2 -> offset limit
	offset := (page - 1) * limit

	// Fetch users using pagination
	err := s.repo.FindWithPagination(
		&users,
		"",                // no specific filter query
		"created_at DESC", // show newest users first
		limit,
		offset,
		[]string{}, // no preloads needed for simple user list
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
