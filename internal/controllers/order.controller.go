package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

// creating the order
func (c *OrderController) CreateOrder(ctx *gin.Context) {

	userID := ctx.GetUint("userID")

	order, err := c.service.CreateOrder(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// Getorders
// GetOrders
func (c *OrderController) GetOrders(ctx *gin.Context) {

	userID := ctx.GetUint("userID")

	// Read pagination params from URL: /orders?page=1&limit=10
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

	// Fetch orders from service with pagination
	orders, err := c.service.GetOrders(userID, page, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {

	userID := ctx.GetUint("userID")

	idParam := ctx.Param("id")
	orderID, _ := strconv.Atoi(idParam)

	order, err := c.service.GetOrderByID(userID, uint(orderID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
