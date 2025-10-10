package s_auth

import (
	"errors"

	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/auth"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type AuthService struct {
	repo *r_auth.AuthRepo
}

func NewAuthService(repo *r_auth.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(username string, password string) (int64, error) {
	createdUserId, err := as.repo.CreateUser(username, utils.GeneratePasswordHash(password))
	if err != nil {
		return 0, err
	}
	return createdUserId, err
}

func (as *AuthService) LoginUser(username string, password string) (int64, error) {
	gotUser, err := as.repo.GetUser(username, utils.GeneratePasswordHash(password))
	if err != nil {
		return 0, err
	}

	if gotUser.Password != utils.GeneratePasswordHash(password) {
		return 0, errors.New("wrong password")
	}

	return gotUser.Id, err
}
