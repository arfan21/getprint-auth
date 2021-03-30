package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type RefreshToken struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	UserID    uuid.UUID      `json:"user_id"`
}
