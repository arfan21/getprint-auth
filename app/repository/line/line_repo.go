package line

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type LineRepository interface {
	VerifyIdToken(ctx context.Context, idToken string) (*LineVerifyIdTokenResponse, error)
}

type lineRepository struct {
	clinetId string
}

func NewLineRepository() LineRepository {
	lineClientID := os.Getenv("LINE_CLIENT_ID")
	return &lineRepository{lineClientID}
}

func (repo lineRepository) VerifyIdToken(ctx context.Context, idToken string) (*LineVerifyIdTokenResponse, error) {
	URL := "https://api.line.me/oauth2/v2.1/verify"
	reqBody := strings.NewReader(fmt.Sprintf("id_token=%s&client_id=%s", idToken, repo.clinetId))

	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, "POST", URL, reqBody)
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

	tes := make(map[string]interface{})
	err = json.Unmarshal(body, &tes)
	lineVerifyIdTokenResponse := new(LineVerifyIdTokenResponse)
	err = json.Unmarshal(body, &lineVerifyIdTokenResponse)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, errors.New("user not found")
	}

	if res.StatusCode >= 404 && res.StatusCode < 500 {
		return nil, errors.New("bad request")
	}

	if res.StatusCode >= 500 {
		return nil, errors.New("internal server error")
	}

	return lineVerifyIdTokenResponse, nil
}
