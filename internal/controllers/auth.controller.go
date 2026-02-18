package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
	
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
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	err := a.authService.Register(
		req.Username,
		req.Email,
		req.Password,
	)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Registration failed", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", nil)
}

// Login function _ controller
func (a *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid login credentials", err.Error())
		return
	}

	tokens, err := a.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
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
	userIDParam := c.Param("userId")

	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", "The provided user ID is not a valid number")
		return
	}

	err = a.authService.SendEmailOTP(uint(userID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to send OTP", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "OTP sent successfully", nil)
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
	userIDParam := c.Param("userId")

	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", "The provided user ID is not a valid number")
		return
	}

	var req struct {
		OTP string `json:"otp" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "OTP is required", err.Error())
		return
	}

	err = a.authService.VerifyEmailOTP(uint(userID), req.OTP)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Verification failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Email verified successfully", nil)
}

// RefreshToken function for refreshing access token
func (a *AuthController) RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Refresh token is required", err.Error())
		return
	}

	tokens, err := a.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid refresh token", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Token refreshed successfully", gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
