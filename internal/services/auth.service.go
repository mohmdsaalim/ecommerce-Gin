package services

import (
	"errors"
	// "fmt"
	// "time"

	// "github.com/mohmdsaalim/ecommerce-Gin/cmd/worker"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
	// "github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	repo repositories.Repository
	redis *redis.Client
}
// costructer from main -> service
func NewAuthService(repo repositories.Repository, redisClient *redis.Client) *AuthService {
	return &AuthService{repo: repo, redis: redisClient}
}

// register service
func (s *AuthService) Register(username, email, password string) error {
	// need to check the user alredy exist...
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

	return  s.repo.Insert(user)

	// otp, _:= utils.GenerateOTP(1)
	// key := fmt.Sprintf("otp:email_verification:%d", user.ID)

	// err = s.redis.Set(database.Ctx, key, otp, 5*time.Minute).Err()
	// if err != nil{
	// 	return err
	// }
	// // worker.OTPChanne
	// return nil
}

// validation 

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
