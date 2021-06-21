package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arfan21/getprint-service-auth/app/controllers/middleware"
	models2 "github.com/arfan21/getprint-service-auth/app/models"
	"github.com/arfan21/getprint-service-auth/app/repository/mysql"
	"github.com/arfan21/getprint-service-auth/app/repository/user"
	"github.com/dgrijalva/jwt-go"
)

type RefreshTokenService interface {
	Create(refreshToken *models2.RefreshToken) error
	UpdateTokenByRefreshToken(refreshToken string) (map[string]interface{}, error)
}

type refreshTokenService struct {
	repo mysql.RefreshTokenRepository
}

func NewRefreshTokenService(repo mysql.RefreshTokenRepository) RefreshTokenService {
	return &refreshTokenService{repo: repo}
}

func (srv refreshTokenService) Create(refreshToken *models2.RefreshToken) error {
	return srv.repo.Create(refreshToken)
}

func (srv refreshTokenService) UpdateTokenByRefreshToken(refreshToken string) (map[string]interface{}, error) {
	data, err := srv.repo.GetByToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid Token")
	}

	oldToken, err := middleware.VerifyToken(data.Token, "refreshToken")
	if err != nil {
		return nil, err
	}
	claims, ok := oldToken.Claims.(jwt.MapClaims)

	if !ok || !oldToken.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	claimsJson, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("Invalid Token")
	}

	userResponse := new(user.UserResoponseData)

	json.Unmarshal(claimsJson, userResponse)

	jwtExp, _ := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)
	jwtExpUnix := time.Now().Add(time.Minute * time.Duration(jwtExp)).Unix()
	newToken, err := middleware.CreateToken(*userResponse, "GetprintIDToken", jwtExpUnix, "token")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": newToken,
	}, nil
}
