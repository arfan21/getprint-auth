package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"service-auth/services"
)

func main() {
	//kid := os.Getenv("RSA_KEY_ID")
	//clientId := os.Getenv("CLIENT_ID")
	//clientSecret := os.Getenv("CLIENT_SECRET")
	//privKey := utils.CreateKey(kid)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	route := echo.New()

	route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Getprint service authentication")
	})
	route.GET("/line-callback", func(c echo.Context) error {
		idToken := c.QueryParam("id_token")
		res, err := services.VerifyIdTokenLine(context.Background(), idToken)
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	})
	route.Static("/oauth", "oauth")

	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
