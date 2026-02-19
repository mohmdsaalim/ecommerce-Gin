package services

import (
	"errors"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type OrderService struct {
	repo repositories.Repository
}

func NewOrderService(repo repositories.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(userID uint) (*models.Order, error) {

	var cart models.Cart

	// 1️⃣ Get Cart with Items
	err := s.repo.FindOne(
		&cart,
		"user_id = ?",
		[]string{"Items"},
		userID,
	)
	if err != nil {
		return nil, errors.New("cart not found")
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	// 2️⃣ Calculate total
	var total float64
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Price
	}

	// 3️⃣ Create Order
	order := models.Order{
		UserID:     userID,
		TotalPrice: total,
		Status:     "pending",
	}

	if err := s.repo.Insert(&order); err != nil {
		return nil, err
	}

	// 4️⃣ Create Order Items
	for _, item := range cart.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			VariantID: item.VariantID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		s.repo.Insert(&orderItem)
	}

	// 5️⃣ Clear Cart Items
	for _, item := range cart.Items {
		s.repo.UpdateFields(&models.Product{}, item.ProductID, map[string]interface{}{"is_carted": false})
	}
	s.repo.Delete(&models.CartItem{}, "cart_id = ?", cart.ID)

	return &order, nil
}

// get orders
// GetOrders retrieves a user's orders with pagination
func (s *OrderService) GetOrders(userID uint, page, limit int) ([]models.Order, error) {
	var orders []models.Order

	// Calculate offset for pagination
	// page 1 -> offset 0
	// page 2 -> offset limit
	offset := (page - 1) * limit

	// Fetch orders using pagination
	err := s.repo.FindWithPagination(
		&orders,
		"user_id = ?",
		"created_at DESC", // Show latest orders first
		limit,
		offset,
		[]string{
			"OrderItems",
			"OrderItems.Product",
			"OrderItems.Variant",
		},
		userID,
	)
	return orders, err
}

func (s *OrderService) GetOrderByID(userID, orderID uint) (*models.Order, error) {
	var order models.Order

	err := s.repo.FindOne(
		&order,
		"id = ? AND user_id = ?",
		[]string{"OrderItems"},
		orderID, userID,
	)

	return &order, err
}
