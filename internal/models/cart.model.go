package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	User      *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items,omitempty"`
}

type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `gorm:"not null;index" json:"cart_id"`
	ProductID uint      `gorm:"not null;index" json:"product_id"`
	VariantID uint      `gorm:"not null;index" json:"variant_id"`
	Quantity  int       `gorm:"not null;default:1" json:"quantity"`
	Price     float64   `gorm:"not null;type:decimal(10,2)" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	Cart      Cart           `gorm:"foreignKey:CartID" json:"cart,omitempty"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Variant   ProductVariant `gorm:"foreignKey:VariantID" json:"variant,omitempty"`
}