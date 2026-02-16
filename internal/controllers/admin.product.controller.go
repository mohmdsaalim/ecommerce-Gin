package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type AdminProductController struct {
	service *services.AdminProductService
}

func NewAdminProductController(service *services.AdminProductService) *AdminProductController {
	return &AdminProductController{service: service}
}

func (c *AdminProductController) CreateProduct(ctx *gin.Context) {

	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *AdminProductController) GetProductByID(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	product, err := c.service.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *AdminProductController) GetAllProducts(ctx *gin.Context) {

	// Read pagination params from URL: /admin/products?page=1&limit=10
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

	// Fetch products from service with pagination
	products, err := c.service.GetAllProducts(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *AdminProductController) UpdateProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	var data map[string]interface{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🛡️ SECURITY: Whitelist only allowed fields to update
	allowedFields := map[string]bool{
		"name":         true,
		"description":  true,
		"base_price":   true,
		"category":     true,
		"sub_category": true,
		"image_url":    true,
		"is_active":    true,
	}

	filteredData := make(map[string]interface{})
	for key, value := range data {
		if allowedFields[key] {
			// Fix: The frontend sends "image_url", but the database column is "primary_image".
			// We map it here to avoid the "column image_url does not exist" error.
			if key == "image_url" {
				filteredData["primary_image"] = value
			} else {
				filteredData[key] = value
			}
		}
	}

	if len(filteredData) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no valid fields to update"})
		return
	}

	if err := c.service.UpdateProduct(uint(id), filteredData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

func (c *AdminProductController) DeleteProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := c.service.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func (c *AdminProductController) UpdateStock(ctx *gin.Context) {

	idParam := ctx.Param("variantId")
	id, _ := strconv.Atoi(idParam)

	var body struct {
		Stock int `json:"stock"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateStock(uint(id), body.Stock); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "stock updated"})
}
