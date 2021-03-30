package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func LoginUser(ctx context.Context, email, password string) (map[string]interface{}, error) {
	url := os.Getenv("SERVICE_USER")
	data := map[string]interface{}{
		"email" : email,
		"password" : password,
	}

	dataJson, err := json.Marshal(data)
	if err != nil{
		return nil, errors.New("internal server error")
	}

	payload := bytes.NewBuffer(dataJson)
	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, "POST", url + "/login", payload)
	req.Header.Add("Content-Type","application/json")
	if err != nil {
		return nil, errors.New("internal server error")
	}

	res, err := client.Do(req)

	if err != nil {
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

	if res.StatusCode == 500 {
		return nil, errors.New(decodeJSON["message"].(string))
	}

	if res.StatusCode == 404 {
		return nil, errors.New(decodeJSON["message"].(string))
	}

	decodeJSON["status_code"] = res.StatusCode
	return decodeJSON, nil
}
