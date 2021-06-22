package http

import (
	"net/http"
	"strings"

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

	cookieRefreshToken := helpers.NewRefreshTokenCookie(dataToken["refresh_token"].(string))
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, helpers.Response("success", nil, map[string]interface{}{
		"token": dataToken["token"].(string),
	}))
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

	cookieRefreshToken := helpers.NewRefreshTokenCookie(dataToken["refresh_token"].(string))
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, helpers.Response("success", nil, map[string]interface{}{
		"token": dataToken["token"].(string),
	}))
}

func (ctrl authController) VerifyToken(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" {
		return c.JSON(http.StatusUnauthorized, helpers.Response("error", "unauthorized", nil))
	}

	token := strings.Split(authorizationHeader, "Bearer ")[1]

	data, err := ctrl.authSrv.VerifyToken(token)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.Response("error", err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (ctrl authController) Logout(c echo.Context) error {
	cookieRefreshToken := helpers.RemoveRefreshTokenCookie()
	c.SetCookie(cookieRefreshToken)

	return c.JSON(http.StatusOK, helpers.Response("success", nil, nil))
}
