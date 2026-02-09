package models

import "time"

type Address struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`

	FullName  string    `gorm:"size:100;not null" json:"full_name"`
	Phone     string    `gorm:"size:20;not null" json:"phone"`
	Address1  string    `gorm:"size:255;not null" json:"address1"`
	Address2  string    `gorm:"size:255" json:"address2"`
	City      string    `gorm:"size:100;not null" json:"city"`
	State     string    `gorm:"size:100;not null" json:"state"`
	Pincode   string    `gorm:"size:20;not null" json:"pincode"`
	Country   string    `gorm:"size:100;default:India" json:"country"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User *User `gorm:"foreignKey:UserID" json:"-"`
}