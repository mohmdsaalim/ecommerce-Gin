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
// func (s *ProductService) GetAllProducts(category, subCategory string) ([]models.Product, error) {

// 	var products []models.Product
// 	var query string
// 	var args []interface{}

// 	switch {
// 	// Both category and subcategory
// 	case category != "" && subCategory != "":
// 		query = "category = ? AND sub_category = ? AND is_active = ?"
// 		args = []interface{}{category, subCategory, true}

// 	// Only category
// 	case category != "":
// 		query = "category = ? AND is_active = ?"
// 		args = []interface{}{category, true}

// 	//  filter
// 	case subCategory != "":
// 		query = "sub_category = ? AND is_active = ?"
// 		args = []interface{}{subCategory, true}
// 	default:
// 		query = "is_active = ?"
// 		args = []interface{}{true}
// 	}

// 	err := s.repo.FindAll(
// 		&products,
// 		query,
// 		"created_at DESC",
// 		[]string{"Variants"},
// 		args...,
// 	)

// 	return products, err
// }


func (s *ProductService) GetAllProducts(category, subCategory, search string) ([]models.Product, error) {

	var products []models.Product
	var query string
	var args []interface{}

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

	err := s.repo.FindAll(
		&products,
		query,
		"created_at DESC",
		[]string{"Variants"},
		args...,
	)

	return products, err
}







func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {

	var product models.Product

	err := s.repo.FindOne(
		&product,
		"id = ? AND is_active = ?",
		id,
		true,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}