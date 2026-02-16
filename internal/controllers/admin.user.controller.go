package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type AdminUserController struct {
	service *services.AdminUserService
}

func NewAdminUserController(service *services.AdminUserService) *AdminUserController {
	return &AdminUserController{service: service}
}

func (c *AdminUserController) GetAllUsers(ctx *gin.Context) {

	// Read pagination params from URL: /admin/users?page=1&limit=10
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	// Convert page to number
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	// Convert limit to number
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Fetch users from service
	users, err := c.service.GetAllUsers(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *AdminUserController) GetUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *AdminUserController) DeactivateUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.DeactivateUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deactivated"})
}

func (c *AdminUserController) ActivateUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.ActivateUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user activated"})
}

type ChangeRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

func (c *AdminUserController) ChangeRole(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var req ChangeRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.ChangeRole(uint(id), req.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "role updated"})
}

func (c *AdminUserController) DeleteUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
