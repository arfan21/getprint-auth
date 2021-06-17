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

func (ctrl refreshTokenControllers) RefreshToken(c echo.Context) error {

	cookie, err := c.Cookie("X-GETPRINT-REFRESH")
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error, nil))
	}

	newToken, err := ctrl.rtSrv.UpdateTokenByRefreshToken(cookie.Value)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", utils.CustomErrors(err), nil))
	}
	cookieToken := new(http.Cookie)
	cookieToken.Name = "X-GETPRINT-KEY"
	cookieToken.Value = newToken["token"].(string)
	cookieToken.MaxAge = 1
	cookieToken.HttpOnly = true
	cookieToken.Secure = true
	cookieToken.Domain = "*.localhost"
	cookieToken.Path = "/"

	return c.JSON(http.StatusOK, utils.Response("success", nil, nil))
}
