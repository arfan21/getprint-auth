package refreshToken

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/echo/v4"
	"net/http"
	"service-auth/utils"
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

func (ctrl refreshTokenControllers) RefreshToken(c echo.Context) error{
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error, nil))
	}

	return nil
}