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
func (s *ProductService) ListProducts() ([]models.Product, error) {
	// Use the specialized method that loads categories and variants
	return s.repo.FindAllProductsWithCategory()
}

func (s *ProductService) GetProductById(id uint) (*models.Product, error) {
	// Use the specialized method that loads categories and variants
	return s.repo.FindProductWithCategory(id)
}

