package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Data    UserResoponseData `json:"data,omitempty"`
}

type UserResoponseData struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Name          string    `json:"name"`
	Picture       string    `json:"picture"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	PhoneNumber   string    `json:"phone_number"`
	Address       string    `json:"address"`
	Role          string    `json:"role"`
	Provider      string    `json:"provider"`
	ProviderID    string    `json:"provider_id"`
	LastLogin     time.Time `json:"last_login"`
}
