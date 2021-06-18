package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	refreshTokenRepo "github.com/arfan21/getprint-service-auth/repository/mysql/refreshToken"
	"github.com/arfan21/getprint-service-auth/services/auth"
	"github.com/arfan21/getprint-service-auth/services/line"
	refreshTokenSrv "github.com/arfan21/getprint-service-auth/services/refreshToken"
	"github.com/arfan21/getprint-service-auth/utils"
)

type AuthController interface {
	Routes(router *echo.Echo)
	VerifyToken(c echo.Context) error
}

type authController struct {
	authSrv auth.AuthService
	lineSrv line.LineService
}

func NewAuthController(db *gorm.DB) AuthController {
	rtRepo := refreshTokenRepo.NewRefreshTokenRepository(db)
	rtSrv := refreshTokenSrv.NewRefreshTokenService(rtRepo)
	authSrv := auth.NewAuthService(rtSrv)
	lineSrv := line.NewLineService(authSrv)
	return &authController{authSrv, lineSrv}
}

func (ctrl authController) Login(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error(), nil))
	}

	dataToken, err := ctrl.authSrv.Login(data["email"].(string), data["password"].(string))

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	cookieToken := new(http.Cookie)
	cookieToken.Name = "getprint-jwt"
	cookieToken.Value = dataToken["token"].(string)
	cookieToken.MaxAge = 0
	cookieToken.HttpOnly = true
	cookieToken.Secure = true
	cookieToken.Path = "/"
	c.SetCookie(cookieToken)
	cookieRefreshToken := new(http.Cookie)
	cookieRefreshToken.Name = "getprint-refresh-token"
	cookieRefreshToken.Value = dataToken["refresh_token"].(string)
	cookieRefreshToken.MaxAge = 0
	cookieRefreshToken.HttpOnly = true
	cookieRefreshToken.Secure = true
	cookieRefreshToken.Path = "/"
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, utils.Response("success", nil, nil))
}

func (ctrl authController) CallbackLine(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error(), nil))
	}

	dataToken, err := ctrl.lineSrv.CallbackHandler(data["id_token"].(string))

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}

	cookieToken := new(http.Cookie)
	cookieToken.Name = "getprint-jwt"
	cookieToken.Value = dataToken["token"].(string)
	cookieToken.MaxAge = 0
	cookieToken.HttpOnly = true
	cookieToken.Secure = true

	cookieToken.Path = "/"
	c.SetCookie(cookieToken)
	cookieRefreshToken := new(http.Cookie)
	cookieRefreshToken.Name = "getprint-refresh-token"
	cookieRefreshToken.Value = dataToken["refresh_token"].(string)
	cookieRefreshToken.MaxAge = 0
	cookieRefreshToken.HttpOnly = true
	cookieRefreshToken.Secure = true

	cookieRefreshToken.Path = "/"
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, utils.Response("success", nil, nil))
}

func (ctrl authController) VerifyToken(c echo.Context) error {
	cookie, err := c.Cookie("getprint-jwt")

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error(), nil))
	}

	data, err := ctrl.authSrv.VerifyToken(cookie.Value)

	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.Response("error", err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.Response("success", nil, data))
}

func (ctrl authController) Logout(c echo.Context) error {
	_, err := c.Cookie("getprint-jwt")

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error(), nil))
	}

	cookieToken := new(http.Cookie)
	cookieToken.Name = "getprint-jwt"
	cookieToken.Value = ""
	cookieToken.MaxAge = -1
	cookieToken.HttpOnly = true
	cookieToken.Secure = true

	cookieToken.Path = "/"
	c.SetCookie(cookieToken)
	cookieRefreshToken := new(http.Cookie)
	cookieRefreshToken.Name = "getprint-refresh-token"
	cookieRefreshToken.Value = ""
	cookieRefreshToken.MaxAge = -1
	cookieRefreshToken.HttpOnly = true
	cookieRefreshToken.Secure = true

	cookieRefreshToken.Path = "/"
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, utils.Response("success", nil, nil))
}
