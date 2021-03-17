package refreshToken

import (
	"gorm.io/gorm"
	"service-auth/models"
)

type RefreshRepository interface {
	Create(refreshToken *models.RefreshToken) error
	GetByToken(token string) (error, *models.RefreshToken)
	DeleteByUserID(id uint) error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshRepository {
	return &refreshTokenRepository{db}
}

func (repo refreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return repo.db.Create(refreshToken).Error
}

func (repo refreshTokenRepository) GetByToken(token string) (error, *models.RefreshToken) {
	refreshToken := new(models.RefreshToken)
	err := repo.db.Where("token = ?", token).Find(&refreshToken).Error
	if err != nil {
		return err , nil
	}
	return nil, refreshToken
}

func (repo refreshTokenRepository) DeleteByUserID(id uint) error {
	return repo.db.Unscoped().Where("id=?", id).Delete(models.RefreshToken{}).Error
}