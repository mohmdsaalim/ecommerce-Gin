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
	cartService *services.CartService,
	) {
	// injecting the service into -> controller 
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	cartController := controllers.NewCartController(cartService)
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
	}

	// User routes 
	user := r.Group("/")
	user.Use(middlewares.AuthMiddleWare())
	{
		user.GET("/profile") // for checking purpose
		user.GET("/cart", cartController.GetCart)// completed -> checked 
		user.POST("/cart/items", cartController.AddToCart)// completed -> checked 
		user.PUT("/cart/item/:id", cartController.UpdateItem)
		user.DELETE("/cart/items/:id", cartController.RemoveItem)// completed -> checked 
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