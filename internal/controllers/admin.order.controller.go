package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type AdminOrderController struct {
	service *services.AdminOrderService
}

func NewAdminOrderController(service *services.AdminOrderService) *AdminOrderController {
	return &AdminOrderController{service: service}
}

func (c *AdminOrderController) GetAllOrders(ctx *gin.Context) {

	orders, err := c.service.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *AdminOrderController) GetOrder(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	order, err := c.service.GetOrderByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func (c *AdminOrderController) UpdateStatus(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var req UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.UpdateStatus(uint(id), req.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "order status updated"})
}


func (c *AdminOrderController) DeleteOrder(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.DeleteOrder(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}