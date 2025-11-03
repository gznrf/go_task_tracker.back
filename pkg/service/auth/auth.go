package s_auth

import (
	"errors"

	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type AuthService struct {
	repo *repo.Repo
}

func NewAuthService(repo *repo.Repo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(data *m_auth.RegisterRequest) (*m_auth.RegisterResponse, error) {
	var output *m_auth.RegisterResponse

	data.Password = utils.GeneratePasswordHash(data.Password)
	output, err := s.repo.AuthRepo.RegisterUser(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *AuthService) LoginUser(data *m_auth.LoginRequest) (*m_auth.LoginResponse, error) {
	var output *m_auth.LoginResponse

	data.Password = utils.GeneratePasswordHash(data.Password)
	output, passFromDb, err := s.repo.AuthRepo.LoginUser(data)
	if err != nil {
		return nil, err
	}

	if passFromDb != data.Password {
		return nil, errors.New("wrong password")
	}

	return output, err
}
