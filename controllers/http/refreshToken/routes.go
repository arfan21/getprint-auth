package refreshToken

import "github.com/labstack/echo/v4"

func (ctrl *refreshTokenControllers) Routes(route *echo.Echo) {
	route.POST("/refresh-token", ctrl.RefreshToken)
}
