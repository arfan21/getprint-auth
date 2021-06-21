package services

import (
	"context"
	"fmt"

	line2 "github.com/arfan21/getprint-service-auth/app/repository/line"
	user2 "github.com/arfan21/getprint-service-auth/app/repository/user"
)

type LineService interface {
	CallbackHandler(idToken string) (map[string]interface{}, error)
}

type lineService struct {
	lineRepo line2.LineRepository
	userRepo user2.UserRepository
	authSrv  AuthService
}

func NewLineService(authSrv AuthService) LineService {
	lineRepo := line2.NewLineRepository()
	userRepo := user2.NewUserRepository(context.Background())
	return &lineService{lineRepo, userRepo, authSrv}
}

func (srv lineService) CallbackHandler(idToken string) (map[string]interface{}, error) {
	dataLine, err := srv.lineRepo.VerifyIdToken(context.Background(), idToken)
	fmt.Println("data line :", dataLine)
	if err != nil {
		return nil, err
	}

	dataUser, err := srv.userRepo.LoginLine(*dataLine)
	fmt.Println("data user :", dataUser)
	if err != nil {
		return nil, err
	}

	return srv.authSrv.Auth(*dataUser)
}
