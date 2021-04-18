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
}

func NewAuthService(refreshTokenSrv _refreshTokenSrv.RefreshTokenService) AuthService {
	return &authService{refreshTokenSrv}
}

func (srv authService) Login(email, password string) (map[string]interface{}, error) {
	user := _userRepo.NewUserRepository(context.Background())
	data, err := user.Login(email, password)
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
	uuidString := data["id"].(string)
	refreshTokenModel := new(models.RefreshToken)
	refreshTokenModel.Email = email
	refreshTokenModel.Token = refreshToken
	refreshTokenModel.UserID = uuid.FromStringOrNil(uuidString)
	err = srv.refreshTokenSrv.Create(refreshTokenModel)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":         token,
		"refresh_token": refreshToken,
	}, nil
}
