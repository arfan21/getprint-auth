package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	refreshTokenRepo "github.com/arfan21/getprint-service-auth/repository/mysql/refreshToken"
	"github.com/arfan21/getprint-service-auth/services/auth"
	refreshTokenSrv "github.com/arfan21/getprint-service-auth/services/refreshToken"
	"github.com/arfan21/getprint-service-auth/utils"
)

type AuthController interface {
	Routes(router *echo.Echo)
}

type authController struct {
	authSrv auth.AuthService
}

func NewAuthController(db *gorm.DB) AuthController {
	rtRepo := refreshTokenRepo.NewRefreshTokenRepository(db)
	rtSrv := refreshTokenSrv.NewRefreshTokenService(rtRepo)
	authSrv := auth.NewAuthService(rtSrv)
	return &authController{authSrv}
}

func (ctrl authController) Login(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error, nil))
	}

	dataToken, err := ctrl.authSrv.Login(data["email"].(string), data["password"].(string))

	if err != nil {

		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", nil, dataToken))
}
