package refreshToken

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/echo/v4"
)

type RefreshTokenControllers interface {
	Routes(route *echo.Echo)
}

type refreshTokenControllers struct {
	srv *server.Server
}

func NewRefreshTokenControllers(srv *server.Server) RefreshTokenControllers {
	return &refreshTokenControllers{srv: srv}
}
