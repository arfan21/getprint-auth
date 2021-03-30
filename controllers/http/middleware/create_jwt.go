package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"service-auth/models"
	"service-auth/utils"
)

var AUD = os.Getenv("AUDIENCE")
var ISS = os.Getenv("ISSUER")

//CreateToken ... typeToken is token or refreshToken
func CreateToken(data map[string]interface{}, kid string, exp int64, typeToken string) (string ,error){
	privKey, _ , err := utils.ReadKey(typeToken)
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, models.JwtClaims{Email: data["email"].(string), StandardClaims: jwt.StandardClaims{
		Audience: AUD,
		Issuer: ISS,
		ExpiresAt: exp,
		Subject: data["id"].(string),
	}})

	return token.SignedString(key)
}