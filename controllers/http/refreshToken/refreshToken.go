package refreshToken

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	refreshTokenRepo "github.com/arfan21/getprint-service-auth/repository/mysql/refreshToken"
	refreshTokenSrv "github.com/arfan21/getprint-service-auth/services/refreshToken"
	"github.com/arfan21/getprint-service-auth/utils"
)

type RefreshTokenControllers interface {
	Routes(route *echo.Echo)
}

type refreshTokenControllers struct {
	rtSrv refreshTokenSrv.RefreshTokenService
}

func NewRefreshTokenControllers(db *gorm.DB) RefreshTokenControllers {
	rtRepo := refreshTokenRepo.NewRefreshTokenRepository(db)
	rtSrv := refreshTokenSrv.NewRefreshTokenService(rtRepo)
	return &refreshTokenControllers{rtSrv}
}

func (ctrl refreshTokenControllers) RefreshToken(c echo.Context) error{
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error, nil))
	}

	newToken, err := ctrl.rtSrv.UpdateTokenByRefreshToken(data["refresh_token"].(string),data["email"].(string))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, newToken))
}