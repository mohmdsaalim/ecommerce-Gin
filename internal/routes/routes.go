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
	orderService *services.OrderService,
	userSrevice *services.UserService,
	) {
	// injecting the service into -> controller 
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	cartController := controllers.NewCartController(cartService)
	orderController := controllers.NewOrderController(orderService)
	userController := controllers.NewUserController(userSrevice)
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
		user.GET("/profile", userController.GetProfile)// completed -> checked
		user.PUT("/profile", userController.UpdateProfile)// completed -> checked 
		user.POST("/profile/address", userController.AddAddress)// completed -> checked 
		user.GET("/cart", cartController.GetCart)// completed -> checked 
		user.POST("/cart/items", cartController.AddToCart)// completed -> checked 
		user.PUT("/cart/item/:id", cartController.UpdateItem)
		user.DELETE("/cart/items/:id", cartController.RemoveItem)// completed -> checked 
		user.POST("/orders", orderController.CreateOrder)// completed -> checked 
		user.GET("/orders", orderController.GetOrders)// completed -> checked 
		user.GET("orders/:id", orderController.GetOrderByID)// completed -> checked 
				// adding payments last
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