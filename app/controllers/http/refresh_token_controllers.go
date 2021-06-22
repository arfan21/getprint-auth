package http

import (
	"net/http"

	"github.com/arfan21/getprint-service-auth/app/helpers"
	"github.com/arfan21/getprint-service-auth/app/services"
	"github.com/labstack/echo/v4"
)

type RefreshTokenControllers interface {
	RefreshToken(c echo.Context) error
}

type refreshTokenControllers struct {
	rtSrv services.RefreshTokenService
}

func NewRefreshTokenControllers(rtSrv services.RefreshTokenService) RefreshTokenControllers {
	return &refreshTokenControllers{rtSrv}
}

func (ctrl refreshTokenControllers) RefreshToken(c echo.Context) error {

	cookie, err := c.Cookie("getprint-refresh-token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}

	newToken, err := ctrl.rtSrv.UpdateTokenByRefreshToken(cookie.Value)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", helpers.CustomErrors(err), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, map[string]interface{}{
		"token": newToken["token"].(string),
	}))
}
