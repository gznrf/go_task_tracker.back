package service

import (
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/auth"
)

type Middleware interface {
}

type Auth interface {
	CreateUser(username string, password string) (int64, error)
	LoginUser(username string, password string) (int64, error)
}

type Service struct {
	AuthService Auth
}

func NewService(authRepo *r_auth.AuthRepo) *Service {
	return &Service{
		AuthService: s_auth.NewAuthService(authRepo),
	}
}
