docker ps ---> list running containers
docker exec -it ecommerce_postgres psql -U ecommerce_user -d ecommerce_db ---> postgres terminal
docker exec -it ecommerce_redis redis-cli ---> redis CLI

imp     docker-compose up -d ---> start project in background 
imp     docker-compose down ---> stop project

docker logs -f ecommerce_postgres ---> follow logs






DOCKER COMMANDS – SINGLE PAGE (WITH USE CASES)

docker --version
# Check Docker is installed

docker info
# Check Docker status and system info

docker ps
# List running containers

docker ps -a
# List all containers (running + stopped)

docker start <container_name_or_id>
# Start a stopped container

docker stop <container_name_or_id>
# Stop a running container

docker restart <container_name_or_id>
# Restart a container

docker rm <container_name_or_id>
# Remove a stopped container

docker rm -f <container_name_or_id>
# Force stop and remove a container

docker rm -f $(docker ps -aq)
# Stop and remove ALL containers (dev cleanup)

docker logs <container_name_or_id>
# View container logs

docker logs -f <container_name_or_id>
# Follow logs live

docker logs --tail=50 <container_name_or_id>
# View last 50 log lines

docker exec -it <container_name_or_id> sh
# Open shell inside container

docker exec -it <container_name_or_id> bash
# Open bash shell (if available)

docker exec -it ecommerce_postgres psql -U ecommerce_user -d ecommerce_db
# Open PostgreSQL terminal inside container

docker exec -it ecommerce_redis redis-cli
# Open Redis CLI inside container

docker images
# List docker images

docker rmi <image_id>
# Remove docker image

docker image prune
# Remove unused images

docker volume ls
# List volumes (DB data stored here)

docker volume inspect <volume_name>
# Inspect volume details

docker volume prune
# Remove unused volumes (deletes DB data)

docker network ls
# List docker networks

docker network inspect <network_name>
# Inspect network details

docker-compose up
# Start project services

docker-compose up -d
# Start project services in background

docker-compose down
# Stop project services

docker-compose down -v
# Stop services and remove volumes (DB reset)

docker-compose up -d --build
# Rebuild and restart services

docker-compose logs
# View logs of all services

docker-compose logs -f
# Follow logs live

docker-compose logs postgres
# View logs of postgres service

docker-compose ps
# List services in compose project

Request
 → Routes
   → Controller
     → Service
       → Repository
         → PostgreSQL




                imp NOTES. daily use
--------------------------------------------
 
# to see the postgresDB inside the docker
 // docker exec -it ecommerce_postgres psql -U ecommerce_user -d ecommerce_db 

# to stop the docker 
 // docker compose down

# start the docker 
 // docker compose up -d

# check the container are running 
 // docker ps



# Injections
	Use PostgreSQL as Repository
	•	Inject it into AuthService
	•	AuthService is injected into Controller
	•	Controller is used by routes



❌ When Generic Repo Is NOT Enough

When you need:
	•	Preload("Category")
	•	Joins()
	•	Complex filtering
	•	Pagination
	•	Sorting
	•	Aggregations
	•	Business logic filtering













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

// NewProductController creates a new product controller instance
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// ============================================
// PRODUCT RETRIEVAL ENDPOINTS
// ============================================

// GetProducts handles multiple query scenarios:
// GET /products - Get all products
// GET /products?category=kits - Get products by category
// GET /products?sub_category=home - Get products by subcategory
// GET /products?search=barcelona - Search products
func (ctrl *ProductController) GetProducts(c *gin.Context) {
	category := c.Query("category")
	subCategory := c.Query("sub_category")
	search := c.Query("search")

	var products []models.Product
	var err error

	// Handle different query parameters
	switch {
	case search != "":
		// Search functionality
		products, err = ctrl.service.SearchProducts(search)
		
	case category != "":
		// Filter by category
		products, err = ctrl.service.GetProductsByCategory(category)
		
	case subCategory != "":
		// Filter by subcategory
		products, err = ctrl.service.GetProductsBySubCategory(subCategory)
		
	default:
		// Get all products
		products, err = ctrl.service.GetAllProducts()
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   len(products),
		"data":    products,
	})
}















// GetProductByID retrieves a single product by ID
// GET /products/:id
func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid product ID",
		})
		return
	}

	product, err := ctrl.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

// ============================================
// PRODUCT MANAGEMENT ENDPOINTS (ADMIN)
// ============================================

// CreateProduct creates a new product
// POST /products
// Body: JSON product data
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := ctrl.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Product created successfully",
		"data":    product,
	})
}

// UpdateProduct updates an existing product
// PUT /products/:id
// Body: JSON with fields to update
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid product ID",
		})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := ctrl.service.UpdateProduct(uint(id), updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product updated successfully",
	})
}

// DeleteProduct soft deletes a product
// DELETE /products/:id
func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid product ID",
		})
		return
	}

	if err := ctrl.service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product deleted successfully",
	})
}

// ============================================
// ADDITIONAL ENDPOINTS
// ============================================

// CheckProductAvailability checks if a product variant has enough stock
// GET /products/check-availability/:variant_id?quantity=2
func (ctrl *ProductController) CheckProductAvailability(c *gin.Context) {
	variantID, err := strconv.ParseUint(c.Param("variant_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid variant ID",
		})
		return
	}

	quantityStr := c.DefaultQuery("quantity", "1")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid quantity",
		})
		return
	}

	isAvailable, err := ctrl.service.CheckProductAvailability(uint(variantID), quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"available": isAvailable,
		"variant_id": variantID,
		"quantity":   quantity,
	})
}