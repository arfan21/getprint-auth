package authControllers

import "github.com/labstack/echo/v4"

func (ctrl authControllers) Token(c echo.Context) error  {
	grantType := c.FormValue("grant_type")
	if grantType == "line_id_token"{

		return c.JSON(200, "tes")
	}
	return ctrl.srv.HandleTokenRequest(c.Response(), c.Request())
}

