package model

import (
	"time"
)

type Tables struct {
	ID           int       `gorm:"primaryKey"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Password     string    `gorm:"column:password"`
	Email        string    `gorm:"column:email"`
	UserType     string    `gorm:"column:user_type"`
	Token        string    `gorm:"column:token"`
	RefreshToken string    `gorm:"column:refresh_token"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	UserId       string    `gorm:"column:user_id"`
}
