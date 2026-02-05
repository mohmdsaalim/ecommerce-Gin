package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// GET /product  public API -> get all product
func(pc *ProductController) GetProducts(c *gin.Context){
	products, err := pc.service.ListProducts()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data":products})
}

//  GET / product/:id pubils API -> get single product
func (pc *ProductController) GetProductById(c *gin.Context){
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid id"})
		return
	}

	product, err := pc.service.GetProductById(uint(id))
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":product})
}


func (pc *ProductController) GetProductsBycat(c *gin.Context) {
	categorySlug := c.Query("category")

	var products []models.Product
	var err error

	if categorySlug != "" {
        // Filter by category
        // products, err = pc.service.GetProductsBycategory(categorySlug)
    } else {
        // Get all products
        products, err = pc.service.ListProducts()
    }
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "count":   len(products),
        "data":    products,
    })
}