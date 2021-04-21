package refreshToken

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/arfan21/getprint-service-auth/controllers/http/middleware"
	"github.com/arfan21/getprint-service-auth/models"
	refreshTokenRepo "github.com/arfan21/getprint-service-auth/repository/mysql/refreshToken"
	_userRepo "github.com/arfan21/getprint-service-auth/repository/user"
)

type RefreshTokenService interface {
	Create(refreshToken *models.RefreshToken) error
	UpdateTokenByRefreshToken(refreshToken, email string) (map[string]interface{}, error)
}

type refreshTokenService struct {
	repo refreshTokenRepo.RefreshTokenRepository
}

func NewRefreshTokenService(repo refreshTokenRepo.RefreshTokenRepository) RefreshTokenService {
	return &refreshTokenService{repo: repo}
}

func (srv refreshTokenService) Create(refreshToken *models.RefreshToken) error {
	return srv.repo.Create(refreshToken)
}

func (srv refreshTokenService) UpdateTokenByRefreshToken(refreshToken, email string) (map[string]interface{}, error) {
	data, err := srv.repo.GetByToken(refreshToken)
	if err != nil {
		return nil, err
	}

	oldToken, err := middleware.VerifyToken(refreshToken, "refreshToken")
	if err != nil {
		return nil, err
	}
	claims, ok := oldToken.Claims.(jwt.MapClaims)

	if !ok || !oldToken.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}
	if claims["email"].(string) != email {
		return nil, fmt.Errorf("Invalid Email")
	}

	dataForNewToken := _userRepo.UserLoginResponse{
		ID:    data.UserID.String(),
		Email: data.Email,
		Role:  data.Role,
	}

	jwtExp, _ := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)
	jwtExpUnix := time.Now().Add(time.Minute * time.Duration(jwtExp)).Unix()
	newToken, err := middleware.CreateToken(dataForNewToken, "GetprintIDToken", jwtExpUnix, "token")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": newToken,
	}, nil
}
