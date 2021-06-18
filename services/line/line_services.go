package line

import (
	"context"
	"fmt"

	"github.com/arfan21/getprint-service-auth/repository/line"
	"github.com/arfan21/getprint-service-auth/repository/user"
	"github.com/arfan21/getprint-service-auth/services/auth"
)

type LineService interface {
	CallbackHandler(idToken string) (map[string]interface{}, error)
}

type lineService struct {
	lineRepo line.LineRepository
	userRepo user.UserRepository
	authSrv  auth.AuthService
}

func NewLineService(authSrv auth.AuthService) LineService {
	lineRepo := line.NewLineRepository()
	userRepo := user.NewUserRepository(context.Background())
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
