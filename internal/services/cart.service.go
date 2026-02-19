package services

import (
	"errors"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/repositories"
)

type CartService struct {
	repo repositories.Repository
}

func NewCartService(repo repositories.Repository) *CartService {
	return &CartService{repo: repo}
}

// Get cart item
func (s *CartService) GetCart(userID uint) (*models.Cart, error) {
	var cart models.Cart

	err := s.repo.FindOne(&cart, "user_id = ?", []string{"Items", "Items.Product", "Items.Variant"}, userID)

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

// Add to Cart
func (s *CartService) AddToCart(userID, productID, variantID uint, qty int) error {

	if qty <= 0 {
		return errors.New("invalid quantity")
	}

	var cart models.Cart

	// 1️Check if cart exists
	err := s.repo.FindOne(&cart, "user_id = ?", nil, userID)

	if err != nil {
		// If not found → create new cart
		cart = models.Cart{
			UserID: userID,
		}
		if err := s.repo.Insert(&cart); err != nil {
			return err
		}
	}

	// 2️ Check if item already exists
	var item models.CartItem
	err = s.repo.FindOne(
		&item,
		"cart_id = ? AND product_id = ? AND variant_id = ?", nil,
		cart.ID, productID, variantID,
	)

	if err == nil {
		// Already exists → return error
		return errors.New("item already in cart")
	}

	// 3️ Insert new cart item
	newItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		VariantID: variantID,
		Quantity:  qty,
		Price:     100, // TODO: fetch actual product price
	}

	if err := s.repo.Insert(&newItem); err != nil {
		return err
	}

	// Update product is_carted to true
	return s.repo.UpdateFields(&models.Product{}, productID, map[string]interface{}{"is_carted": true})
}

// update cart item
func (s *CartService) UpdateCartItem(itemID uint, qty int) error {

	if qty <= 0 {
		return errors.New("invalid quantity")
	}

	var item models.CartItem
	if err := s.repo.FindByID(&item, itemID); err != nil {
		return err
	}

	item.Quantity = qty

	return s.repo.UpdateByID(&models.CartItem{}, itemID, item)
}

// to remove the cart item
func (s *CartService) RemoveItem(itemID uint) error {
	var item models.CartItem
	if err := s.repo.FindByID(&item, itemID); err != nil {
		return err
	}

	if err := s.repo.Delete(&models.CartItem{}, "id = ?", itemID); err != nil {
		return err
	}

	// Update product is_carted to false
	return s.repo.UpdateFields(&models.Product{}, item.ProductID, map[string]interface{}{"is_carted": false})
}
