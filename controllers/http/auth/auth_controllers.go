package auth

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	refreshTokenRepo "service-auth/repository/mysql/refreshToken"
	"service-auth/services/auth"
	refreshTokenSrv "service-auth/services/refreshToken"
	"service-auth/utils"
)

type AuthController interface {
}

type authController struct {
	authSrv auth.AuthService
}

func NewAuthController(db *gorm.DB) AuthController {
	rtRepo := refreshTokenRepo.NewRefreshTokenRepository(db)
	rtSrv := refreshTokenSrv.NewRefreshTokenService(rtRepo)
	authSrv := auth.NewAuthService(rtSrv)
	return &authController{authSrv}
}

func (ctrl authController) Login(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response("error", err.Error, nil))
	}

	dataToken, err := ctrl.authSrv.Login(data["email"].(string), data["password"].(string))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response("error", err.Error, nil))
	}

	return c.JSON(http.StatusOK, utils.Response("success", "", dataToken))
}
