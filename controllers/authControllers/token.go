package authControllers

import "github.com/labstack/echo/v4"

func (ctrl authControllers) Token(c echo.Context) error  {
	return ctrl.srv.HandleTokenRequest(c.Response(), c.Request())
}

