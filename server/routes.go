package server

import (
	"os"
	"path/filepath"

	"github.com/arfan21/getprint-service-auth/app/controllers/http"
	"github.com/arfan21/getprint-service-auth/app/repository/mysql"
	"github.com/arfan21/getprint-service-auth/app/services"
	"github.com/arfan21/getprint-service-auth/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(mysqlClient config.Client) *echo.Echo{
	route := echo.New()
	route.Use(middleware.Recover())
	route.Use(middleware.Logger())

	path, _ := os.Getwd()
	wellKnowsRoute := route.Group("/.well-knows")
	wellKnowsRoute.Use(middleware.Static(filepath.Join(path+"/well-knows")))

	refreshTokenRepo := mysql.NewRefreshTokenRepository(mysqlClient)
	refreshTokenSrv := services.NewRefreshTokenService(refreshTokenRepo)
	refreshTokenCtrl := http.NewRefreshTokenControllers(refreshTokenSrv)

	authSrv := services.NewAuthService(refreshTokenSrv)
	lineSrv := services.NewLineService(authSrv)
	authCtrl := http.NewAuthController(authSrv, lineSrv)

	apiV1 := route.Group("/v1")

	// routing auth
	authRoute := apiV1.Group("/auth")
	authRoute.POST("/login", authCtrl.Login)
	authRoute.POST("/verify", authCtrl.VerifyToken)
	authRoute.POST("/line-callback", authCtrl.CallbackLine)
	authRoute.POST("/logout", authCtrl.Logout)
	authRoute.POST("/refresh-token", refreshTokenCtrl.RefreshToken)

	return route
}