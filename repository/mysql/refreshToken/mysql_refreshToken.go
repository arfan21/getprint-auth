package refreshToken

import (
	"gorm.io/gorm"

	"github.com/arfan21/getprint-service-auth/models"
)

type RefreshTokenRepository interface {
	Create(refreshToken *models.RefreshToken) error
	GetByToken(token string) (*models.RefreshToken, error)
	DeleteByUserID(id uint) error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (repo refreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return repo.db.Create(refreshToken).Error
}

func (repo refreshTokenRepository) GetByToken(token string) (*models.RefreshToken, error){
	refreshToken := new(models.RefreshToken)
	err := repo.db.Where("token = ?", token).Find(&refreshToken).Error
	if err != nil {
		return nil , err
	}
	return refreshToken, nil
}

func (repo refreshTokenRepository) DeleteByUserID(id uint) error {
	return repo.db.Unscoped().Where("id=?", id).Delete(models.RefreshToken{}).Error
}