package services

import (
	"errors"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type AdminOrderService struct {
	repo repositories.Repository
}

func NewAdminOrderService(repo repositories.Repository) *AdminOrderService {
	return &AdminOrderService{repo: repo}
}

// GetAllOrders retrieves all orders with pagination for the admin panel
func (s *AdminOrderService) GetAllOrders(page, limit int) ([]models.Order, error) {

	var orders []models.Order

	// Calculate offset for pagination
	// page 1 -> offset 0
	// page 2 -> offset limit
	offset := (page - 1) * limit

	// Fetch orders using pagination
	err := s.repo.FindWithPagination(
		&orders,
		"",                // no specific filter query
		"created_at DESC", // show newest orders first
		limit,
		offset,
		[]string{
			"User",
			"OrderItems",
			"OrderItems.Product",
			"OrderItems.Variant",
		},
	)

	return orders, err
}

func (s *AdminOrderService) GetOrderByID(id uint) (*models.Order, error) {

	var order models.Order

	err := s.repo.FindOne(
		&order,
		"id = ?",
		[]string{
			"User",
			"OrderItems",
			"OrderItems.Product",
			"OrderItems.Variant",
		},
		id,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *AdminOrderService) UpdateStatus(id uint, status string) error {

	validStatuses := map[string]bool{
		"pending":   true,
		"paid":      true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	return s.repo.UpdateFields(
		&models.Order{},
		id,
		map[string]interface{}{
			"status": status,
		},
	)
}

func (s *AdminOrderService) DeleteOrder(id uint) error {

	return s.repo.Delete(
		&models.Order{},
		"id = ?",
		id,
	)
}
