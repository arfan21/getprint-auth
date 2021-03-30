package refreshToken

import (
	"service-auth/models"
	refreshTokenRepo "service-auth/repository/mysql/refreshToken"
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

func (srv refreshTokenService) UpdateTokenByRefreshToken(refreshToken, email string) error {

	return nil
}