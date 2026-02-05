package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;size:255" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"not null;type:decimal(10,2)" json:"price"`
	Category    string         `gorm:"not null;size:50;index" json:"category"`       // "kits" or "lifestyle"
	SubCategory string         `gorm:"size:50;index" json:"sub_category"`            // "home", "away", "goalkeeper", "hoodie", "retro", "casuals"
	Stock       int            `gorm:"default:0" json:"stock"`
	Size        string         `gorm:"size:10" json:"size"`                          // "S", "M", "L", "XL", "XXL"
	ImageURL    string         `gorm:"size:500" json:"image_url"`
	IsAvailable bool           `gorm:"default:true;index" json:"is_available"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}