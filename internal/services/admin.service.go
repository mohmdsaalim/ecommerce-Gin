package services

import (
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type AdminService struct {
	repo repositories.Repository
}

func NewAdminService(repo repositories.Repository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) GetDashboard() (*models.AdminDashboard, error) {

	totalUsers, err := s.repo.Count(&models.User{}, "")
	if err != nil {
		return nil, err
	}

	totalOrders, err := s.repo.Count(&models.Order{}, "")
	if err != nil {
		return nil, err
	}

	totalProducts, err := s.repo.Count(&models.Product{}, "")
	if err != nil {
		return nil, err
	}

	totalRevenue, err := s.repo.Sum(&models.Order{}, "total_price", "")
	if err != nil {
		return nil, err
	}

	pendingOrders, err := s.repo.Count(&models.Order{}, "status = ?", "pending")
	if err != nil {
		return nil, err
	}

	var recentOrders []models.Order
	err = s.repo.FindWithLimit(
		&recentOrders,
		"",
		"created_at DESC",
		5,
		[]string{"OrderItems"},
	)
	if err != nil {
		return nil, err
	}

	dashboard := models.AdminDashboard{
		TotalUsers:    totalUsers,
		TotalOrders:   totalOrders,
		TotalProducts: totalProducts,
		TotalRevenue:  totalRevenue,
		PendingOrders: pendingOrders,
		RecentOrders:  recentOrders,
	}

	return &dashboard, nil
}
