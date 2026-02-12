package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/controllers"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/middlewares"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/services"
)

func RegisterRoute(
	r *gin.Engine,
	authService *services.AuthService,
	productService *services.ProductService,
	cartService *services.CartService,
	orderService *services.OrderService,
	userSrevice *services.UserService,
	adminService *services.AdminService,
	adminProductService *services.AdminProductService,
	adminUserService *services.AdminUserService,
	adminOrderService *services.AdminOrderService,
) {
	// injecting the service into -> controller
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	cartController := controllers.NewCartController(cartService)
	orderController := controllers.NewOrderController(orderService)
	userController := controllers.NewUserController(userSrevice)
	adminController := controllers.NewAdminController(adminService)
	adminProductController := controllers.NewAdminProductController(adminProductService)
	adminUserController := controllers.NewAdminUserController(adminUserService)
	adminOrderController := controllers.NewAdminOrderController(adminOrderService)
	// checking route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// PUBLIC routes
	auth := r.Group("/")
	{
		auth.POST("/auth/register", authController.Register)        // completed -> checked
		auth.GET("/request-email-otp/:userId")                      // send OTP
		auth.POST("/verify-email-otp/:userId")                      // verify OTP
		auth.POST("/auth/login", authController.Login)              // completed -> checked
		auth.GET("/products", productController.GetProducts)        // completed -> checked
		auth.GET("/products/:id", productController.GetProductByID) // completed -> checked
	}

	// User routes
	user := r.Group("/")
	user.Use(middlewares.AuthMiddleWare())
	{
		user.GET("/profile", userController.GetProfile)          // completed -> checked
		user.PUT("/profile", userController.UpdateProfile)       // completed -> checked
		user.POST("/profile/address", userController.AddAddress) // completed -> checked
		user.GET("/cart", cartController.GetCart)                // completed -> checked
		user.POST("/cart/items", cartController.AddToCart)       // completed -> checked
		user.PUT("/cart/item/:id", cartController.UpdateItem)
		user.DELETE("/cart/items/:id", cartController.RemoveItem) // completed -> checked
		user.POST("/orders", orderController.CreateOrder)         // completed -> checked
		user.GET("/orders", orderController.GetOrders)            // completed -> checked
		user.GET("orders/:id", orderController.GetOrderByID)      // completed -> checked
		// adding payments last
	}

	// admin routes
	admin := r.Group("/admin")
	admin.Use(
		middlewares.AuthMiddleWare(),
		middlewares.RequareRole("admin"),
	)
	{
		// admin dashboard
		admin.GET("/dashboard", adminController.GetDashboard)

		// admin prodcts page
		admin.POST("/products", adminProductController.CreateProduct)
		admin.GET("/products", adminProductController.GetAllProducts)
		admin.GET("/products/:id", adminProductController.GetProductByID)
		admin.PUT("/products/:id", adminProductController.UpdateProduct)
		admin.DELETE("/products/:id", adminProductController.DeleteProduct)
		admin.PUT("/variants/:variantId/stock", adminProductController.UpdateStock)

		// admin user cntrls
		admin.GET("/users", adminUserController.GetAllUsers)
		admin.GET("/users/:id", adminUserController.GetUser)
		admin.PUT("/users/:id/deactivate", adminUserController.DeactivateUser)
		admin.PUT("/users/:id/activate", adminUserController.ActivateUser)
		admin.PUT("/users/:id/role", adminUserController.ChangeRole)
		admin.DELETE("/users/:id", adminUserController.DeleteUser)


		// admin orders
		admin.GET("/orders", adminOrderController.GetAllOrders)
		admin.GET("/orders/:id", adminOrderController.GetOrder)
		admin.PUT("/orders/:id/status", adminOrderController.UpdateStatus)
		admin.DELETE("/orders/:id", adminOrderController.DeleteOrder)
	}
}
