package models

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	UserID    uint      `json:"user_id"`
}
