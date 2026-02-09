package models

import "time"

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	ProductID uint      `gorm:"not null"`
	VariantID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product
	Variant ProductVariant
}