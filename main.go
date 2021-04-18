package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"

	_authCtrl "github.com/arfan21/getprint-service-auth/controllers/http/auth"
	_refreshTokenCtrl "github.com/arfan21/getprint-service-auth/controllers/http/refreshToken"
	"github.com/arfan21/getprint-service-auth/utils"
)

func main() {
	db, err := utils.Connect()
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

	route.Static("/well-knows", "oauth")
	authCtrl := _authCtrl.NewAuthController(db)
	authCtrl.Routes(route)

	rtCtrl := _refreshTokenCtrl.NewRefreshTokenControllers(db)
	rtCtrl.Routes(route)

	if err := utils.CreateKey("GetprintIDToken", "token"); err != nil {
		log.Println(err)
	}
	if err := utils.CreateKey("GetprintRefreshToken", "refreshToken"); err != nil {
		log.Println(err)
	}
	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
