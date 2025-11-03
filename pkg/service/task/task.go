package s_task

import (
	"github.com/gznrf/go_task_tracker.back.git/models/task"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
)

type TaskService struct {
	repo *repo.Repo
}

func NewTaskService(repo *repo.Repo) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(data *m_task.CreateRequest) (*m_task.CreateResponse, error) {
	var output *m_task.CreateResponse

	output, err := s.repo.TaskRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *TaskService) Get() (*m_task.GetResponse, error) {
	var output *m_task.GetResponse

	output, err := s.repo.TaskRepo.Get()
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *TaskService) GetByColumnId(data *m_task.GetByColumnIdRequest) (*m_task.GetByColumnIdResponse, error) {
	var output *m_task.GetByColumnIdResponse

	output, err := s.repo.TaskRepo.GetByColumnId(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *TaskService) GetById(data *m_task.GetByIdRequest) (*m_task.GetByIdResponse, error) {
	var output *m_task.GetByIdResponse

	output, err := s.repo.TaskRepo.GetById(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *TaskService) Update(data *m_task.UpdateRequest) (*m_task.UpdateResponse, error) {
	var output *m_task.UpdateResponse

	output, err := s.repo.TaskRepo.Update(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *TaskService) Delete(data *m_task.DeleteRequest) (*m_task.DeleteResponse, error) {
	var output *m_task.DeleteResponse

	output, err := s.repo.TaskRepo.Delete(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
