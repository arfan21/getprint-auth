package user

type UserResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Data    UserLoginResponse `json:"data,omitempty"`
}

type UserLoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
