package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/controllers"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/middlewares"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)


func RegisterRoute(
	r *gin.Engine ,
	authService *services.AuthService,
	productService *services.ProductService,
	) {
	// injecting the service into -> controller 
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)

	// checking route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status":"OK"})
	})

	// PUBLIC routes
	auth := r.Group("/")
	{
		auth.POST("/auth/register", authController.Register) // completed -> checked
		auth.POST("/auth/login", authController.Login) // completed -> checked
		auth.GET("/products", productController.GetProducts)// completed -> checked
		auth.GET("/products/:id",productController.GetProductByID )// completed -> checked 
		auth.GET("/products/kits",)// completed -> checked today 06
		auth.GET("/products/lifestyles",)// completed -> checked today 06

	}

	// User routes 
	user := r.Group("/")
	user.Use(middlewares.AuthMiddleWare())
	{
		user.GET("/profile") // for checking purpose
		user.GET("/cart")
		user.POST("/cart/items")
		user.PUT("/cart/item/:id")
		user.DELETE("/cart/items/:id")
		user.POST("/orders")
		user.GET("/orders")
		user.GET("orders/:id")
	}

	// admin routes
	admin := r.Group("/admin")
	admin.Use(
		middlewares.AuthMiddleWare(),
		middlewares.RequareRole("ADMIN"),
	)
	{

	}
}