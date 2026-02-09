package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}


// get the profile
func (uc *UserController) GetProfile(c *gin.Context) {

	userIDValue, _ := c.Get("userID")
	userID := userIDValue.(uint)

	user, err := uc.service.GetProfile(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}
// Update Post
func (uc *UserController) UpdateProfile(c *gin.Context) {

	userIDValue, _ := c.Get("userID")
	userID := userIDValue.(uint)

	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.UpdateProfile(userID, input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Profile updated successfully"})
}

// post the address
func (uc *UserController) AddAddress(c *gin.Context) {

	userIDValue, _ := c.Get("userID")
	userID := userIDValue.(uint)

	var address models.Address

	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.AddAddress(userID, &address); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Address added successfully"})
}