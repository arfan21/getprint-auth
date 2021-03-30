package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	_authCtrl "service-auth/controllers/http/auth"
	"service-auth/services"
	"service-auth/utils"
)

func main() {
	db, err :=utils.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

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
	route.Static("/well-knows", "oauth")
	authCtrl := _authCtrl.NewAuthController(db)
	authCtrl.Routes(route)

	if err := utils.CreateKey("GetprintIDToken", "token"); err != nil{
		log.Println(err)
	}
	if err := utils.CreateKey("GetprintRefreshToken", "refreshToken"); err != nil{
		log.Println(err)
	}
	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
