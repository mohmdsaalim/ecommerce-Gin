package models

import (
	"time"
)

type Order struct {
	ID         uint         `gorm:"primaryKey"`
	UserID     uint         `gorm:"not null"`
	TotalPrice float64      `gorm:"not null"`
	Status     string       `gorm:"default:'pending'"` // pending, paid, shipped, delivered
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User       User
	OrderItems []OrderItem
}