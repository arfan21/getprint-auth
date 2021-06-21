package http

import (
	"net/http"

	"github.com/arfan21/getprint-service-auth/app/helpers"
	"github.com/arfan21/getprint-service-auth/app/services"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(c echo.Context) error
	CallbackLine(c echo.Context) error
	VerifyToken(c echo.Context) error
	Logout(c echo.Context) error
}

type authController struct {
	authSrv services.AuthService
	lineSrv services.LineService
}

func NewAuthController(authSrv services.AuthService, lineSrv services.LineService) AuthController {
	return &authController{authSrv, lineSrv}
}

func (ctrl authController) Login(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}

	dataToken, err := ctrl.authSrv.Login(data["email"].(string), data["password"].(string))

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
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

	return c.JSON(http.StatusOK, helpers.Response("success", nil, nil))
}

func (ctrl authController) CallbackLine(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}

	dataToken, err := ctrl.lineSrv.CallbackHandler(data["id_token"].(string))

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
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

	return c.JSON(http.StatusOK, helpers.Response("success", nil, nil))
}

func (ctrl authController) VerifyToken(c echo.Context) error {
	cookie, err := c.Cookie("getprint-jwt")

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}

	data, err := ctrl.authSrv.VerifyToken(cookie.Value)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (ctrl authController) Logout(c echo.Context) error {
	_, err := c.Cookie("getprint-jwt")

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
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

	return c.JSON(http.StatusOK, helpers.Response("success", nil, nil))
}
