package models

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct{
	Email string `json:"email"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}
