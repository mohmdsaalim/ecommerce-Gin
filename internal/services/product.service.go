package services

import (
	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type ProductService struct {
	repo repositories.Repository
}

func NewProductService(repo repositories.Repository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}


// GetAllProduct - returns all active products
func (s *ProductService) GetAllProducts(category string) ([]models.Product, error) {
	var products []models.Product

	query := "is_active = ?"
	args := []interface{}{true}

	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}

	err := s.repo.FindAll(
		&products,
		query,
		"created_at DESC",
		[]string{"Variants"},
		args...,
	)

	return products, err
}