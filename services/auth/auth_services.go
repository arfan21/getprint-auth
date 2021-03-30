package auth

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"service-auth/controllers/http/middleware"
	"service-auth/models"
	"service-auth/services"
	"service-auth/services/refreshToken"
	"strconv"
	"time"
)

type AuthService interface {
	Login(email, password string) (map[string]interface{}, error)
}

type authService struct {
	refreshTokenSrv refreshToken.RefreshTokenService
}

func NewAuthService(refreshTokenSrv refreshToken.RefreshTokenService) AuthService {
	return &authService{refreshTokenSrv}
}

func (srv authService) Login(email, password string) (map[string]interface{}, error) {
	data, err := services.LoginUser(context.Background(), email, password)
	if err != nil {
		return nil, err
	}

	jwtExp, _ := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)
	jwtExpUnix := time.Now().Add(time.Minute * time.Duration(jwtExp)).Unix()
	token, err := middleware.CreateToken(data, "GetprintIDToken", jwtExpUnix, "token")

	if err != nil {
		return nil, err
	}

	refreshTokenExp, _ := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXP"), 10, 64)
	refreshTokenExpUnix := time.Now().Add(time.Minute * time.Duration(refreshTokenExp)).Unix()
	refreshToken, err := middleware.CreateToken(data, "GetprintRefreshToken", refreshTokenExpUnix, "refreshToken")

	if err != nil {
		return nil, err
	}

	refreshTokenModel := new(models.RefreshToken)
	refreshTokenModel.Email = email
	refreshTokenModel.Token = refreshToken
	refreshTokenModel.UserID = data["user_id"].(uint)
	err = srv.refreshTokenSrv.Create(refreshTokenModel)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":         token,
		"refresh_token": refreshToken,
	}, nil
}
