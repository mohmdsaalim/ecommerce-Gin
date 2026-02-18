package models

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;size:255" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Category    string         `gorm:"not null;size:50;index" json:"category"`
	SubCategory string         `gorm:"size:50;index" json:"sub_category"`
	BasePrice   float64        `gorm:"not null;type:decimal(10,2)" json:"base_price"`
	SKU         string         `gorm:"unique;size:100" json:"sku"`
	Season      string         `gorm:"size:20" json:"season"`
	
	PrimaryImage   string      `gorm:"size:500" json:"primary_image"`
	SecondaryImage string      `gorm:"size:500" json:"secondary_image"`
	ThumbnailImage string      `gorm:"size:500" json:"thumbnail_image"`
	
	IsActive    bool           `gorm:"default:true;index" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	
	Variants    []ProductVariant `gorm:"foreignKey:ProductID" json:"variants,omitempty"`
}
// product analys ///////
type ProductVariant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null;index" json:"product_id"`
	Size      string    `gorm:"not null;size:10" json:"size"`
	Stock     int       `gorm:"default:0" json:"stock"`
	SKU       string    `gorm:"unique;size:100" json:"sku"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}