package models

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Picture string   `json:"picture"`
	Roles   []string `json:"roles"`
	jwt.StandardClaims
}
