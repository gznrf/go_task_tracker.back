package service

import (
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/user"
)

type Auth interface {
	CreateUser(username string, password string) (int64, error)
	LoginUser(username string, password string) (int64, error)
}
type Middleware interface {
}

type User interface {
	Get() ([]m_user.User, error)
	GetById(userId int64) (m_user.User, error)
	Update(inputData *m_user.UpdateUserRequest) error
	Delete(userId int64) error
}

type Service struct {
	AuthService Auth
	UserService User
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		AuthService: s_auth.NewAuthService(repo),
		UserService: s_user.NewUserService(repo),
	}
}
