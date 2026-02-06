package controllers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/mohmdsaalim/ecommerce-Gin/internal/models"
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

// func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
// 	// category := c.Query("category")
// 	// subcategory := c.Query("sub_category")
// 	// search := c.Query("search")


// 	var products []models.Product
// 	var err error

// 	// handling query params 

// 	// switch{
// 	// case search != "":
// 	// 	products, err = 
// 	// }
// 	products, err = ctrl.service.GetAllProducts()

// 	if err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success":false,
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success":true,
// 		"count":len(products),
// 		"data":products,
// 	})
// }

func (pc *ProductController) GetProducts(c *gin.Context) {
	category := c.Query("category")

	products, err := pc.service.GetAllProducts(category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}