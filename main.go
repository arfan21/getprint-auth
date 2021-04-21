package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authCtrl "github.com/arfan21/getprint-service-auth/controllers/http/auth"
	_refreshTokenCtrl "github.com/arfan21/getprint-service-auth/controllers/http/refreshToken"
	"github.com/arfan21/getprint-service-auth/utils"
)

func main() {
	kidToken := os.Getenv("JWT_KID")
	kidRefreshToken := os.Getenv("REFRESH_TOKEN_KID")

	if err := utils.CreateKey(kidToken, "token"); err != nil {
		log.Println(err)
	}
	if err := utils.CreateKey(kidRefreshToken, "refreshToken"); err != nil {
		log.Println(err)
	}

	db, err := utils.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	route := echo.New()

	route.Use(middleware.Recover())
	route.Use(middleware.Logger())

	route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Getpring Service Auth",
		})
	})

	route.Static("/.well-knows", "well-knows")
	authCtrl := _authCtrl.NewAuthController(db)
	authCtrl.Routes(route)

	rtCtrl := _refreshTokenCtrl.NewRefreshTokenControllers(db)
	rtCtrl.Routes(route)

	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
