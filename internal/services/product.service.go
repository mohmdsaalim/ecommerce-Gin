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

// is in carted set when the item is in cart
// GetAllProducts retrieves products with filtering and pagination
func (s *ProductService) GetAllProducts(category, subCategory, search string, page, limit int) ([]models.Product, error) {

	var products []models.Product
	var query string
	var args []interface{}

	// Basic filtering logic
	switch {
	// Category + SubCategory + Search
	case category != "" && subCategory != "" && search != "":
		query = "category = ? AND sub_category = ? AND LOWER(name) LIKE LOWER(?) AND is_active = ?"
		args = []interface{}{category, subCategory, "%" + search + "%", true}

	// Category + SubCategory
	case category != "" && subCategory != "":
		query = "category = ? AND sub_category = ? AND is_active = ?"
		args = []interface{}{category, subCategory, true}

	// Category + Search
	case category != "" && search != "":
		query = "category = ? AND LOWER(name) LIKE LOWER(?) AND is_active = ?"
		args = []interface{}{category, "%" + search + "%", true}

	// SubCategory + Search
	case subCategory != "" && search != "":
		query = "sub_category = ? AND LOWER(name) LIKE LOWER(?) AND is_active = ?"
		args = []interface{}{subCategory, "%" + search + "%", true}

	// Only Category
	case category != "":
		query = "category = ? AND is_active = ?"
		args = []interface{}{category, true}

	// Only SubCategory
	case subCategory != "":
		query = "sub_category = ? AND is_active = ?"
		args = []interface{}{subCategory, true}

	// Only Search
	case search != "":
		query = "LOWER(name) LIKE LOWER(?) AND is_active = ?"
		args = []interface{}{"%" + search + "%", true}

	// Default (All Active)
	default:
		query = "is_active = ?"
		args = []interface{}{true}
	}

	// Calculate offset: how many items to skip
	// If page is 1 and limit is 10, offset is (1-1)*10 = 0
	// If page is 2 and limit is 10, offset is (2-1)*10 = 1 skip 10
	offset := (page - 1) * limit

	// Use our new FindWithPagination repository method
	err := s.repo.FindWithPagination(
		&products,
		query,
		"created_at DESC",
		limit,
		offset,
		[]string{"Variants"},
		args...,
	)

	return products, err
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {

	var product models.Product

	err := s.repo.FindOne(
		&product,
		"id = ? AND is_active = ?", nil,
		id,
		true,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
