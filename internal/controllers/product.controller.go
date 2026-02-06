package controllers

import (
	// "net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// GetProducts handles multiple query scenarios:
// GET /products - Get all products
// GET /products?category=kits - Get products by category
// GET /products?sub_category=home - Get products by subcategory
// GET /products?search=barcelona - Search products

func (pc *ProductController) GetProducts(c *gin.Context) {

	category := c.Query("category")
	subCategory := c.Query("sub_category")
	search := c.Query("search")

	products, err := pc.service.GetAllProducts(category, subCategory, search)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

// /products/:id
func (pc *ProductController) GetProductByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	product, err := pc.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	c.JSON(200, product)
}