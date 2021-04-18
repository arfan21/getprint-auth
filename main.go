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
	// route.GET("/line-callback", func(c echo.Context) error {
	// 	code := c.QueryParam("code")
	// 	res, err := services.VerifyIdTokenLine(context.Background(), idToken)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(res)
	// 	fmt.Println(code)

	// 	urlLine := "https://api.line.me/oauth2/v2.1/token"
	// 	lineClientID := os.Getenv("LINE_CLIENT_ID")
	// 	lineSecretID := os.Getenv("LINE_SECRET_ID")
	// 	requestBody := map[string]interface{}{
	// 		"grant_type" : "authorization_code",
	// 		"code" : code,
	// 		"client_id": lineClientID,
	// 		"client_secret" : lineSecretID,
	// 		"redirect_uri" : "https://910f089421a9.ngrok.io/",
	// 	}

	// 	jsonByte, _ := json.Marshal(requestBody)
	// 	reqBody := strings.NewReader(`grant_type=authorization_code&code=`+code+`&client_id=`+lineClientID+`&client_secret=`+lineSecretID+`&redirect_uri=https://620e4cee05cd.ngrok.io/`)

	// 	client := new(http.Client)
	// 	req, err := http.NewRequest("POST", urlLine, reqBody)
	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 	res, err := client.Do(req)

	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}

	// 	defer res.Body.Close()
	// 	body, err := ioutil.ReadAll(res.Body)

	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}

	// 	decodeJSON := make(map[string]interface{})

	// 	err = json.Unmarshal(body, &decodeJSON)

	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}

	// 	fmt.Println(decodeJSON)
	// 	return nil
	// })
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
