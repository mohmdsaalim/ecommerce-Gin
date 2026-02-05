package controllers

import (
	"net/http"

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
func (a *AuthController) Register(c *gin.Context){
	var req models.RegisterRequest // <- from user models
// json -> struct
	if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
			return
	}
	
	err := a.authService.Register(
		req.Username,
		req.Email,
		req.Password,
	)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"user registered successfully"})
}

// Login function _ controller
func (a *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest // <- from models
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}

	token, err := a.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}