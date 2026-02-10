package services

import (
	"errors"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
)

type AuthService struct {
	repo repositories.Repository
}
// costructer from main -> service
func NewAuthService(repo repositories.Repository) *AuthService {
	return &AuthService{repo: repo}
}

// register service
func (s *AuthService) Register(username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil{
		return err
	}

	user := &models.User{
		Username:      username,
		Email:         email,
		PasswordHash:  hashedPassword,
		Role:          "USER",
		Status:        "active",
		EmailVerified: false,
	}

	return s.repo.Insert(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	//  Fetch user by email
	if err := s.repo.FindOne(&user, "email = ?",nil, email,); err != nil {
		return "", errors.New("invalid Email")
	}

	// Check password
	if !utils.Checkpassword(user.PasswordHash, password) {
		return "", errors.New("invalid Password")
	}

	//  Generate JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
