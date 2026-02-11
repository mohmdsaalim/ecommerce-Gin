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
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error(), 
		"code":401,
		"status":false, // need to set everywhere
	})

		return
	}

	token, err := a.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}
// salman review
// validation
// constants 







// func (a *AuthController) RequestEmailOTP(c *gin.Context) {

// 	userIDParam := c.Param("userId")

// 	userID, err := strconv.ParseUint(userIDParam, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
// 		return
// 	}

// 	err = a.authService.SendEmailOTP(uint(userID))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
// }