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

func (s *AdminOrderService) GetAllOrders() ([]models.Order, error) {

	var orders []models.Order

	err := s.repo.FindAll(
		&orders,
		"",
		"created_at DESC",
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