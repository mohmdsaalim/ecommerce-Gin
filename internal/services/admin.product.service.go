package services

import (
	"errors"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type AdminProductService struct {
	repo repositories.Repository
}

func NewAdminProductService(repo repositories.Repository) *AdminProductService {
	return &AdminProductService{repo: repo}
}

func (s *AdminProductService) CreateProduct(product *models.Product) error {

	if product.Name == "" || product.BasePrice <= 0 {
		return errors.New("invalid product data")
	}

	return s.repo.Insert(product)
}

// GetAllProducts retrieves all products with pagination for the admin panel
func (s *AdminProductService) GetAllProducts(page, limit int) ([]models.Product, error) {

	var products []models.Product

	// Calculate offset for pagination
	// page 1 -> offset 0
	// page 2 -> offset limit
	offset := (page - 1) * limit

	// Fetch products using pagination
	err := s.repo.FindWithPagination(
		&products,
		"",                // no specific filter query
		"created_at DESC", // show newest products first
		limit,
		offset,
		[]string{"Variants"},
	)

	return products, err
}

func (s *AdminProductService) GetProductByID(id uint) (*models.Product, error) {

	var product models.Product

	err := s.repo.FindOne(
		&product,
		"id = ?",
		[]string{"Variants"},
		id,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *AdminProductService) UpdateProduct(id uint, data map[string]interface{}) error {
	return s.repo.UpdateFields(&models.Product{}, id, data)
}

func (s *AdminProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(&models.Product{}, "id = ?", id)
}

func (s *AdminProductService) UpdateStock(variantID uint, stock int) error {

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	data := map[string]interface{}{
		"stock": stock,
	}

	return s.repo.UpdateFields(&models.ProductVariant{}, variantID, data)
}
