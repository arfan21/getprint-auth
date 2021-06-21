package middleware

import (
	"os"

	models2 "github.com/arfan21/getprint-service-auth/app/models"
	"github.com/arfan21/getprint-service-auth/app/repository/user"
	"github.com/arfan21/getprint-service-auth/config"
	"github.com/dgrijalva/jwt-go"
)

//CreateToken ... typeToken is token or refreshToken
func CreateToken(user user.UserResoponseData, kid string, exp int64, typeToken string) (string, error) {
	var AUD = os.Getenv("AUDIENCE")
	var ISS = os.Getenv("ISSUER")
	privKey, _, err := config.ReadKey(typeToken)
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, models2.JwtClaims{Name: user.Name, Email: user.Email, Picture: user.Picture, Roles: []string{user.Role}, StandardClaims: jwt.StandardClaims{
		Audience:  AUD,
		Issuer:    ISS,
		ExpiresAt: exp,
		Subject:   user.ID.String(),
	}})
	token.Header["kid"] = kid
	return token.SignedString(key)
}
