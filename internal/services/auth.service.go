package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services/workers"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	repo  repositories.Repository
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
	if err != nil {
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

	err = s.repo.Insert(user)
	if err != nil {
		return err
	}

	// Generate 6 digit OTP
	otp, _ := utils.GenerateOTP(6)
	key := fmt.Sprintf("otp:%s", user.Email)

	// Store OTP in redis for 5 minutes
	err = database.SetOTP(key, otp, 5*time.Minute)
	if err != nil {
		return err
	}

	// Send to worker channel
	workers.OTPChannel <- workers.OTPJob{
		Email: user.Email,
		Code:  otp,
	}

	return nil
}

// validation

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	//  Fetch user by email
	if err := s.repo.FindOne(&user, "email = ?", nil, email); err != nil {
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


// package services

// import (
// 	"errors"
// 	"fmt"
// 	"time"

// 	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
// 	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
// 	"github.com/mohmdsaalim/ecommerce-Gin/internal/services/workers"
// 	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
// 	"github.com/mohmdsaalim/ecommerce-Gin/pkg/database"
// 	"github.com/redis/go-redis/v9"
// )

// type AuthService struct {
// 	repo  repositories.Repository
// 	redis *redis.Client
// }

// // costructer from main -> service
// func NewAuthService(repo repositories.Repository, redisClient *redis.Client) *AuthService {
// 	return &AuthService{repo: repo, redis: redisClient}
// }

// // register service
// func (s *AuthService) Register(username, email, password string) error {
// 	// need to check the user alredy exist...
// 	hashedPassword, err := utils.HashPassword(password)
// 	if err != nil {
// 		return err
// 	}

// 	user := &models.User{
// 		Username:      username,
// 		Email:         email,
// 		PasswordHash:  hashedPassword,
// 		Role:          "USER",
// 		Status:        "active",
// 		EmailVerified: false,
// 	}

// 	err = s.repo.Insert(user)
// 	if err != nil {
// 		return err
// 	}

// 	// Generate 6 digit OTP
// 	otp, _ := utils.GenerateOTP(6)
// 	key := fmt.Sprintf("otp:%s", user.Email)

// 	// Store OTP in redis for 5 minutes
// 	err = database.SetOTP(key, otp, 5*time.Minute)
// 	if err != nil {
// 		return err
// 	}

// 	// Send to worker channel
// 	workers.OTPChannel <- workers.OTPJob{
// 		Email: user.Email,
// 		Code:  otp,
// 	}

// 	return nil
// }

// // validation

// func (s *AuthService) Login(email, password string) (string, error) {
// 	var user models.User

// 	//  Fetch user by email
// 	if err := s.repo.FindOne(&user, "email = ?", nil, email); err != nil {
// 		return "", errors.New("invalid Email")
// 	}

// 	// Check password
// 	if !utils.Checkpassword(user.PasswordHash, password) {
// 		return "", errors.New("invalid Password")
// 	}

// 	//  Generate JWT
// 	token, err := utils.GenerateToken(user.ID, user.Role)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }