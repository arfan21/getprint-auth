package authControllers

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/echo/v4"
)

type AuthControllers interface {
	Auth(c echo.Context) error
	Authorize(c echo.Context) error
	Login(c echo.Context) error
	Token(c echo.Context) error
}

type authControllers struct {
	srv *server.Server
}

func NewAuthControllers(route *echo.Echo, srv *server.Server){
	ctrl := authControllers{srv: srv}
	route.POST("/auth", ctrl.Auth)
	route.POST("/token", ctrl.Token)
	route.POST("/authorize", ctrl.Authorize)
	route.POST("/login", ctrl.Login)
}