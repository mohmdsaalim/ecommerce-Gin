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

// Login service 
func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	if err := s.repo.FindByID(&user, email); err != nil{
		return "", errors.New("invalid credentials")
	}

	if !utils.Checkpassword(user.PasswordHash, password) {
		return "", errors.New("invalid credentails")
	}
	return utils.GenerateToken(user.ID, user.Role)
}