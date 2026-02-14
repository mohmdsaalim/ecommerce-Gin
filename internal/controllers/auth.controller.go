package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

// this is jst typ no data-here
type AuthController struct {
	authService *services.AuthService
}

// constructor func
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register func
func (a *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest // <- from user models
	// json -> struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err := a.authService.Register(
		req.Username,
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

// Login function _ controller
func (a *AuthController) Login(c *gin.Context) {

	var req models.LoginRequest // <- from models
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(),
			"code":   401,
			"status": false, // need to set everywhere
		})

		return
	}

	token, err := a.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// salman review
// validation
// constants

// RequestEmailOTP handles the request to send OTP to user's email
// Route: POST /request-email-otp/:userId
// This function:
// 1. Gets userId from URL parameter
// 2. Calls service to generate OTP and store in Redis for 5 minutes
// 3. Sends OTP to user's email
func (a *AuthController) RequestEmailOTP(c *gin.Context) {
	// Step 1: Get userId from URL parameter (example: /request-email-otp/123)
	userIDParam := c.Param("userId")

	// Step 2: Convert userId from string to number
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Step 3: Call service to:
	// - Generate 6-digit OTP
	// - Store OTP in Redis with 5-minute expiration
	// - Send OTP to user's email
	err = a.authService.SendEmailOTP(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 4: Return success response
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// VerifyEmailOTP handles OTP verification from user input
// Route: POST /verify-email-otp/:userId
// Request body: {"otp": "123456"}
// This function:
// 1. Gets userId from URL parameter
// 2. Gets OTP code from request body
// 3. Calls service to verify OTP from Redis
// 4. If OTP matches, sets User.EmailVerified = true
func (a *AuthController) VerifyEmailOTP(c *gin.Context) {
	// Step 1: Get userId from URL parameter (example: /verify-email-otp/123)
	userIDParam := c.Param("userId")

	// Step 2: Convert userId from string to number
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Step 3: Define structure to receive OTP from request body
	// Request body format: {"otp": "123456"}
	var req struct {
		OTP string `json:"otp" binding:"required"` // OTP is required field
	}

	// Step 4: Parse JSON request body and bind to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 5: Call service to:
	// - Get OTP from Redis using user's email
	// - Compare Redis OTP with user's input (req.OTP)
	// - If they match, set User.EmailVerified = true in database
	// - Delete OTP from Redis after verification
	err = a.authService.VerifyEmailOTP(uint(userID), req.OTP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 6: Return success response - User is now verified!
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
