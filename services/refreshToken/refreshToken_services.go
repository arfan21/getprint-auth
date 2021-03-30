package refreshToken

import (
	"fmt"
	"os"
	"service-auth/controllers/http/middleware"
	"service-auth/models"
	refreshTokenRepo "service-auth/repository/mysql/refreshToken"
	"strconv"
	"time"
)

type RefreshTokenService interface {
	Create(refreshToken *models.RefreshToken) error
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
	claims, ok := oldToken.Claims.(models.JwtClaims)

	if !ok || !oldToken.Valid{
		return nil, fmt.Errorf("Invalid Token")
	}

	if claims.Email != email {
		return nil, fmt.Errorf("Invalid Email")
	}

	dataForNewToken := map[string]interface{}{
		"user_id" : data.UserID,
		"email" : email,
	}

	jwtExp, _ := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)
	jwtExpUnix := time.Now().Add(time.Minute * time.Duration(jwtExp)).Unix()
	newToken, err := middleware.CreateToken(dataForNewToken, "GetprintIDToken", jwtExpUnix, "token")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token" : newToken,
	}, nil
}