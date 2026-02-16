package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type AdminController struct {
	service *services.AdminService
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{service: service}
}

// get admin dashboard
func (c *AdminController) GetDashboard(ctx *gin.Context) {

	dashboard, err := c.service.GetDashboard()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dashboard)
}