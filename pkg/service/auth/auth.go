package s_auth

import (
	"errors"

	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type AuthService struct {
	repo *repo.Repo
}

func NewAuthService(repo *repo.Repo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(username string, password string) (int64, error) {
	createdUserId, err := s.repo.AuthRepo.CreateUser(username, utils.GeneratePasswordHash(password))
	if err != nil {
		return 0, err
	}

	return createdUserId, err
}

func (s *AuthService) LoginUser(username string, password string) (int64, error) {
	gotUser, err := s.repo.AuthRepo.GetUser(username, utils.GeneratePasswordHash(password))
	if err != nil {
		return 0, err
	}

	if gotUser.Password != utils.GeneratePasswordHash(password) {
		return 0, errors.New("wrong password")
	}

	return gotUser.Id, err
}
