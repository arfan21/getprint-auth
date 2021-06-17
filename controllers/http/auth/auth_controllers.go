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

	cookieToken := new(http.Cookie)
	cookieToken.Name = "X-GETPRINT-KEY"
	cookieToken.Value = dataToken["token"].(string)
	cookieToken.MaxAge = 1
	cookieToken.HttpOnly = true
	cookieToken.Secure = true
	cookieToken.Domain = "*.localhost"
	cookieToken.Path = "/"
	c.SetCookie(cookieToken)
	cookieRefreshToken := new(http.Cookie)
	cookieRefreshToken.Name = "X-GETPRINT-REFRESH"
	cookieRefreshToken.Value = dataToken["refresh_token"].(string)
	cookieRefreshToken.MaxAge = 1
	cookieRefreshToken.HttpOnly = true
	cookieRefreshToken.Secure = true
	cookieRefreshToken.Domain = "*.localhost"
	cookieRefreshToken.Path = "/"
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, utils.Response("success", nil, dataToken))
}
