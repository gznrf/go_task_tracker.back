package s_user

import (
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type UserService struct {
	repo *repo.Repo
}

func NewUserService(repo *repo.Repo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Get() (*m_user.GetResponse, error) {
	var output *m_user.GetResponse

	output, err := s.repo.UserRepo.Get()
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *UserService) GetById(data *m_user.GetByIdRequest) (*m_user.GetByIdResponse, error) {
	var output *m_user.GetByIdResponse

	output, err := s.repo.UserRepo.GetById(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *UserService) Update(data *m_user.UpdateRequest) (*m_user.UpdateResponse, error) {
	var output *m_user.UpdateResponse

	data.Password = utils.GeneratePasswordHash(data.Password)
	output, err := s.repo.UserRepo.Update(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *UserService) Delete(data *m_user.DeleteRequest) (*m_user.DeleteResponse, error) {
	var output *m_user.DeleteResponse

	output, err := s.repo.UserRepo.Delete(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
