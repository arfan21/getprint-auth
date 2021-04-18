package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/arfan21/getprint-service-auth/models"
	"github.com/arfan21/getprint-service-auth/utils"
)

//CreateToken ... typeToken is token or refreshToken
func CreateToken(data map[string]interface{}, kid string, exp int64, typeToken string) (string ,error){
	var AUD = os.Getenv("AUDIENCE")
	var ISS = os.Getenv("ISSUER")
	privKey, _ , err := utils.ReadKey(typeToken)
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, models.JwtClaims{Email: data["email"].(string), Roles: []string{data["role"].(string)}, StandardClaims: jwt.StandardClaims{
		Audience: AUD,
		Issuer: ISS,
		ExpiresAt: exp,
		Subject: data["id"].(string),
	}})
	token.Header["kid"] = kid
	return token.SignedString(key)
}