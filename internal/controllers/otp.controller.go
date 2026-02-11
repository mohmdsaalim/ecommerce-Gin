package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

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