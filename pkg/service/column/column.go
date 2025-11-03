package s_column

import (
	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
)

type ColumnService struct {
	repo *repo.Repo
}

func NewColumnService(repo *repo.Repo) *ColumnService {
	return &ColumnService{repo: repo}
}

func (s *ColumnService) Create(data *m_column.CreateRequest) (*m_column.CreateResponse, error) {
	var output *m_column.CreateResponse

	output, err := s.repo.ColumnRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ColumnService) Get() (*m_column.GetResponse, error) {
	var output *m_column.GetResponse

	output, err := s.repo.ColumnRepo.Get()
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ColumnService) GetById(data *m_column.GetByIdRequest) (*m_column.GetByIdResponse, error) {
	var output *m_column.GetByIdResponse

	output, err := s.repo.ColumnRepo.GetById(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ColumnService) GetByBoardId(data *m_column.GetByBoardIdRequest) (*m_column.GetByBoardIdResponse, error) {
	var output *m_column.GetByBoardIdResponse

	output, err := s.repo.ColumnRepo.GetByBoardId(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ColumnService) Update(data *m_column.UpdateRequest) (*m_column.UpdateResponse, error) {
	var output *m_column.UpdateResponse

	output, err := s.repo.ColumnRepo.Update(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ColumnService) Delete(data *m_column.DeleteRequest) (*m_column.DeleteResponse, error) {
	var output *m_column.DeleteResponse

	output, err := s.repo.ColumnRepo.Delete(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
