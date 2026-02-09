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
func (c *OrderController) GetOrders(ctx *gin.Context) {

	userID := ctx.GetUint("userID")

	orders, err := c.service.GetOrders(userID)
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