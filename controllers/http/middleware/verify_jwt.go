package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"github.com/arfan21/getprint-service-auth/utils"
)

func VerifyToken(tokenString, typeToken string) (*jwt.Token, error){
	_, pubKey , err := utils.ReadKey(typeToken)
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
