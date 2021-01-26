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

func LoginUser(ctx context.Context, email, password string) (string, error) {
	url := os.Getenv("SERVICE_USER")
	data := map[string]interface{}{
		"email" : email,
		"password" : password,
	}

	dataJson, err := json.Marshal(data)
	if err != nil{
		return "", errors.New("internal server error")
	}

	payload := bytes.NewBuffer(dataJson)
	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, "GET", url + "user/login", payload)

	if err != nil {
		return "", errors.New("internal server error")
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
		return "", errors.New("internal server error")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	decodeJSON := make(map[string]interface{})

	err = json.Unmarshal(body, &decodeJSON)

	if err != nil {
		return "", err
	}

	if res.StatusCode == 404 {
		return "", errors.New("user not found")
	}

	decodeJSON["status_code"] = res.StatusCode

	return decodeJSON["sub"].(string), nil
}
