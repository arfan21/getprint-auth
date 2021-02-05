package services

import "github.com/labstack/echo/v4"

func LineHandler(c echo.Context) error{

	r := c.Request()
	r.Form.Set("grant_type", "password")
	return nil
}
