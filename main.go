package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"service-auth/controllers/authControllers"
	"service-auth/services"
)

func main(){
	manager := manage.NewDefaultManager()
	//token memeory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	privKey := ReadKey()

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", privKey, jwt.SigningMethodRS256))
	//client memory store
	clientstore := store.NewClientStore()
	clientstore.Set("Z2V0cHJpbnQtc2VydmljZS1hdXRo", &models.Client{
		ID : "Z2V0cHJpbnQtc2VydmljZS1hdXRo",
		Secret: "c2VjcmV0LWtleS1mb3ItZ2V0cHJpbnQtc2VydmljZS1hdXRo",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientstore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	srv.SetPasswordAuthorizationHandler(services.PasswordAuthorizationHandler)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8888"
	}
	route := echo.New()

	route.GET("/", func(c echo.Context) error{
		return c.JSON(http.StatusOK, "Getprint service authentication")
	})
	route.Static("/oauth", "oauth")

	authControllers.NewAuthControllers(route, srv)

	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}

func ReadKey()[]byte{
	privKey, err := ioutil.ReadFile("./key/private.pem")
	if err != nil{
		log.Println(err.Error())
		return nil
	}
	return privKey
}