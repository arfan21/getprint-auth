package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	uuid "github.com/satori/go.uuid"

	"github.com/arfan21/getprint-service-auth/controllers/http/middleware"
	"github.com/arfan21/getprint-service-auth/models"
	_userRepo "github.com/arfan21/getprint-service-auth/repository/user"
	_refreshTokenSrv "github.com/arfan21/getprint-service-auth/services/refreshToken"
)

type AuthService interface {
	Login(email, password string) (map[string]interface{}, error)
}

type authService struct {
	refreshTokenSrv _refreshTokenSrv.RefreshTokenService
	userRepo        _userRepo.UserRepository
}

func NewAuthService(refreshTokenSrv _refreshTokenSrv.RefreshTokenService) AuthService {
	userRepo := _userRepo.NewUserRepository(context.Background())
	return &authService{refreshTokenSrv, userRepo}
}

func (srv authService) Login(email, password string) (map[string]interface{}, error) {
	kidToken := os.Getenv("JWT_KID")
	kidRefreshToken := os.Getenv("REFRESH_TOKEN_KID")
	data, err := srv.userRepo.Login(email, password)
	if err != nil {
		return nil, err
	}

	jwtExp, _ := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)
	jwtExpUnix := time.Now().Add(time.Minute * time.Duration(jwtExp)).Unix()
	token, err := middleware.CreateToken(*data, kidToken, jwtExpUnix, "token")
	if err != nil {
		return nil, err
	}

	refreshTokenExp, _ := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXP"), 10, 64)
	refreshTokenExpUnix := time.Now().Add(time.Minute * time.Duration(refreshTokenExp)).Unix()
	refreshToken, err := middleware.CreateToken(*data, kidRefreshToken, refreshTokenExpUnix, "refreshToken")

	if err != nil {
		return nil, err
	}

	refreshTokenModel := new(models.RefreshToken)
	refreshTokenModel.Token = refreshToken
	refreshTokenModel.Email = data.Email
	refreshTokenModel.UserID = uuid.FromStringOrNil(data.ID)
	refreshTokenModel.Role = data.Role

	err = srv.refreshTokenSrv.Create(refreshTokenModel)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":         token,
		"refresh_token": refreshToken,
	}, nil
}
