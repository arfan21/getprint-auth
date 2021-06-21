package middleware

import (
	"fmt"

	"github.com/arfan21/getprint-service-auth/config"
	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString, typeToken string) (*jwt.Token, error){
	_, pubKey , err := config.ReadKey(typeToken)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func (token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	return token ,err
}
