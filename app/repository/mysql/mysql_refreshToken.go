package mysql

import (
	models2 "github.com/arfan21/getprint-service-auth/app/models"
	"github.com/arfan21/getprint-service-auth/config"
)

type RefreshTokenRepository interface {
	Create(refreshToken *models2.RefreshToken) error
	GetByToken(token string) (*models2.RefreshToken, error)
	DeleteByUserID(id uint) error
}

type refreshTokenRepository struct {
	DB config.Client
}

func NewRefreshTokenRepository(mysqlClient config.Client) RefreshTokenRepository {
	return &refreshTokenRepository{mysqlClient}
}

func (repo refreshTokenRepository) Create(refreshToken *models2.RefreshToken) error {
	return repo.DB.Conn().Create(refreshToken).Error
}

func (repo refreshTokenRepository) GetByToken(token string) (*models2.RefreshToken, error){
	refreshToken := new(models2.RefreshToken)
	err := repo.DB.Conn().Where("token = ?", token).Find(&refreshToken).Error
	if err != nil {
		return nil , err
	}
	return refreshToken, nil
}

func (repo refreshTokenRepository) DeleteByUserID(id uint) error {
	return repo.DB.Conn().Unscoped().Where("id=?", id).Delete(models2.RefreshToken{}).Error
}