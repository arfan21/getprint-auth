package auth

import "github.com/labstack/echo/v4"

func (ctrl authController) Routes(router *echo.Echo){
	router.POST("/login", ctrl.Login)
}