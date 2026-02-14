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

	// ========== OTP GENERATION AND STORAGE IN REDIS ==========
	// Step 1: Generate a random 6-digit OTP (One-Time Password)
	// Example: "123456"
	otp, _ := utils.GenerateOTP(6)

	// Step 2: Create a unique key for storing OTP in Redis
	// Format: "otp:user@email.com"
	// This key is used to retrieve the OTP later during verification
	key := fmt.Sprintf("otp:%s", user.Email)

	// Step 3: Store OTP in Redis with 5-minute expiration (TTL = Time To Live)
	// - Redis will automatically delete the OTP after 5 minutes
	// - This prevents old OTPs from being used after 5 minutes
	// - Redis config (host, port, password, db) is loaded from config.yaml
	err = database.SetOTP(key, otp, 5*time.Minute)
	if err != nil {
		return err
	}

	// Step 4: Send OTP to user's email using worker channel
	// Worker will send email in background without blocking the request
	workers.OTPChannel <- workers.OTPJob{
		Email: user.Email, // User's email address
		Code:  otp,        // The 6-digit OTP to send
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

// SendEmailOTP resends the OTP to the user
func (s *AuthService) SendEmailOTP(userID uint) error {
	var user models.User
	// Find user
	if err := s.repo.FindByID(&user, userID); err != nil {
		return errors.New("user not found")
	}

	// ========== RESEND OTP TO USER ==========
	// Step 1: Generate a new random 6-digit OTP
	// Example: "987654"
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return err
	}

	// Step 2: Create unique Redis key using user's email
	// Format: "otp:user@email.com"
	key := fmt.Sprintf("otp:%s", user.Email)

	// Step 3: Store new OTP in Redis with 5-minute auto-expiration
	// - If old OTP exists, it will be overwritten with new OTP
	// - After 5 minutes, Redis automatically deletes the OTP
	// - Redis connection details come from config.yaml file
	err = database.SetOTP(key, otp, 5*time.Minute)
	if err != nil {
		return err
	}

	// Step 4: Send new OTP to user's email via worker
	workers.OTPChannel <- workers.OTPJob{
		Email: user.Email, // Send to this email
		Code:  otp,        // New 6-digit OTP
	}

	return nil
}

// VerifyEmailOTP verifies the OTP code entered by user
// This function:
// 1. Gets OTP from Redis using user's email
// 2. Compares Redis OTP with user's input
// 3. If they match, sets User.EmailVerified = true in database
// 4. Deletes OTP from Redis to prevent reuse
func (s *AuthService) VerifyEmailOTP(userID uint, code string) error {
	var user models.User

	// Step 1: Find the user in database by ID
	if err := s.repo.FindByID(&user, userID); err != nil {
		return errors.New("user not found")
	}

	// Step 2: Create Redis key using user's email
	// This is the same key format used when storing OTP
	// Format: "otp:user@email.com"
	key := fmt.Sprintf("otp:%s", user.Email)

	// Step 3: Get the OTP stored in Redis
	// - If OTP expired (5 minutes passed), Redis will return error
	// - If OTP doesn't exist, Redis will return error
	storedOTP, err := database.GetOTP(key)
	if err != nil {
		// OTP not found or expired (5 minutes passed)
		return errors.New("OTP expired or invalid")
	}

	// Step 4: Compare OTP from Redis with user's input
	// - storedOTP: OTP saved in Redis when it was generated
	// - code: OTP entered by user in the request
	if storedOTP != code {
		// OTPs don't match - user entered wrong OTP
		return errors.New("invalid OTP")
	}

	// ========== OTP IS CORRECT - UPDATE USER ==========
	// Step 5: Set User.EmailVerified to true in database
	// This marks that user has verified their email successfully
	if !user.EmailVerified {
		// Update only the email_verified field in database
		err = s.repo.UpdateFields(&user, userID, map[string]interface{}{
			"email_verified": true, // Set EmailVerified to true
		})
		if err != nil {
			return err
		}
	}

	// Step 6: Delete OTP from Redis after successful verification
	// This prevents the same OTP from being used again (security)
	// Ignore any error from deletion (using _ to discard error)
	_ = database.DeleteOTP(key)

	// Success! User is now verified
	return nil
}
