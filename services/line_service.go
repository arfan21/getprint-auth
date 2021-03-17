package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func VerifyIdTokenLine(ctx context.Context, idToken string) (map[string]interface{}, error){
	urlLine := "https://api.line.me/oauth2/v2.1/verify"
	lineClientID := os.Getenv("LINE_CLIENT_ID")
	reqBody := strings.NewReader(`id_token=`+idToken+`&client_id=`+lineClientID)

	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, "POST", urlLine, reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, errors.New("internal server error")
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
		return nil, errors.New("internal server error")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	decodeJSON := make(map[string]interface{})

	err = json.Unmarshal(body, &decodeJSON)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, errors.New("user not found")
	}

	decodeJSON["status_code"] = res.StatusCode

	return decodeJSON, nil
}
