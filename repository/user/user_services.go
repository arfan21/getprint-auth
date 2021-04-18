package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type UserRepository interface {
	Login(email, password string) (map[string]interface{}, error)
}

type userRepository struct {
	ctx context.Context
}

func NewUserRepository(ctx context.Context) UserRepository {
	return &userRepository{ctx}
}

func (repo userRepository) Login(email, password string) (map[string]interface{}, error) {
	url := os.Getenv("SERVICE_USER")
	data := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(dataJson)
	client := new(http.Client)
	req, err := http.NewRequestWithContext(repo.ctx, "POST", url+"/login", payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	decodedJSON := make(map[string]interface{})

	err = json.Unmarshal(body, &decodedJSON)

	if err != nil {
		return nil, err
	}

	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return nil, errors.New(decodedJSON["message"].(string))
	}

	return decodedJSON["data"].(map[string]interface{}), nil
}
