package s_project

import (
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
)

type ProjectService struct {
	repo *repo.Repo
}

func NewProjectService(repo *repo.Repo) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) Create(data *m_project.CreateRequest) (*m_project.CreateResponse, error) {
	var output *m_project.CreateResponse

	output, err := s.repo.ProjectRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ProjectService) Get() (*m_project.GetResponse, error) {
	var output *m_project.GetResponse

	output, err := s.repo.ProjectRepo.Get()
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ProjectService) GetByUserId(data *m_project.GetByUserIdRequest) (*m_project.GetByUserIdResponse, error) {
	var output *m_project.GetByUserIdResponse

	output, err := s.repo.ProjectRepo.GetByUserId(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ProjectService) GetById(data *m_project.GetByIdRequest) (*m_project.GetByIdResponse, error) {
	var output *m_project.GetByIdResponse

	output, err := s.repo.ProjectRepo.GetById(data)
	if err != nil {
		return nil, nil
	}
	return output, nil
}
func (s *ProjectService) Update(data *m_project.UpdateRequest) (*m_project.UpdateResponse, error) {
	var output *m_project.UpdateResponse

	output, err := s.repo.ProjectRepo.Update(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ProjectService) Delete(data *m_project.DeleteRequest) (*m_project.DeleteResponse, error) {
	var output *m_project.DeleteResponse

	output, err := s.repo.ProjectRepo.Delete(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
