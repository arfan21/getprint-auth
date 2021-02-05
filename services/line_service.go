package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func VerifyIdTokenLine(ctx context.Context, idToken string) (map[string]interface{}, error){
	urlLine := "https://api.line.me/oauth2/v2.1/verify"
	lineClientID := os.Getenv("LINE_CLIENT_ID")
	data := map[string]interface{}{
		"id_token" : idToken,
		"client_id" : lineClientID,
	}

	dataJson, err := json.Marshal(data)
	if err != nil{
		return nil, errors.New("internal server error")
	}

	payload := bytes.NewBuffer(dataJson)
	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, "POST", urlLine, payload)

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
