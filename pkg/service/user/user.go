package s_user

import (
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
)

type UserService struct {
	repo *repo.Repo
}

func NewUserService(repo *repo.Repo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Get() ([]m_user.User, error) {
	usersList, err := s.repo.UserRepo.Get()
	if err != nil {
		return nil, err
	}
	return usersList, nil
}
func (s *UserService) GetById(userId int64) (m_user.User, error) {
	user, err := s.repo.UserRepo.GetById(userId)
	if err != nil {
		return m_user.User{}, err
	}

	return user, nil
}
func (s *UserService) Update(inputData *m_user.UpdateUserRequest) error {
	err := s.repo.UserRepo.Update(inputData)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserService) Delete(userId int64) error {
	err := s.repo.UserRepo.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
