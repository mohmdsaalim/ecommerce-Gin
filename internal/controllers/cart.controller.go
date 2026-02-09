package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type CartController struct {
	service *services.CartService
}

func NewCartController(service *services.CartService) *CartController {
	return &CartController{service: service}
}

// Get the Cart item 
func (cc *CartController) GetCart(c *gin.Context) {

	userID := c.GetUint("userID") // from JWT middleware

	cart, err := cc.service.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// Add to Cart
type AddCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	VariantID uint `json:"variant_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

func (cc *CartController) AddToCart(c *gin.Context) {

	userID := c.GetUint("userID")

	var req AddCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.service.AddToCart(
		userID,
		req.ProductID,
		req.VariantID,
		req.Quantity,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item added to cart"})
}

// Update cart item
func (cc *CartController) UpdateItem(c *gin.Context) {

	idParam := c.Param("id")
	itemID, _ := strconv.Atoi(idParam)

	var body struct {
		Quantity int `json:"quantity"`
	}

	c.ShouldBindJSON(&body)

	err := cc.service.UpdateCartItem(uint(itemID), body.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}



// Remove item from cart
func (cc *CartController) RemoveItem(c *gin.Context) {

	idParam := c.Param("id")
	itemID, _ := strconv.Atoi(idParam)

	err := cc.service.RemoveItem(uint(itemID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item removed"})
}