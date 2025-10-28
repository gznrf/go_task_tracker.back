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

func (s *ColumnService) Create(creatingColumn *m_column.CreateRequest) (int64, error) {
	createdId, err := s.repo.ColumnRepo.Create(creatingColumn)
	if err != nil {
		return 0, err
	}
	return createdId, nil
}
func (s *ColumnService) Get() ([]m_column.Column, error) {
	columnsList, err := s.repo.ColumnRepo.Get()
	if err != nil {
		return nil, err
	}
	return columnsList, nil
}
func (s *ColumnService) GetById(columnId int64) (m_column.GetByIdResponse, error) {
	outputColumn, err := s.repo.ColumnRepo.GetById(columnId)
	if err != nil {
		return m_column.GetByIdResponse{}, err
	}

	return outputColumn, nil
}
func (s *ColumnService) GetByBoardId(boardId int64) ([]m_column.Column, error) {
	columnsList, err := s.repo.ColumnRepo.GetByBoardId(boardId)
	if err != nil {
		return nil, err
	}

	return columnsList, nil
}
func (s *ColumnService) Update(updatingColumn *m_column.UpdateRequest) error {
	err := s.repo.ColumnRepo.Update(updatingColumn)
	if err != nil {
		return err
	}

	return nil
}
func (s *ColumnService) Delete(deletingColumn *m_column.DeleteRequest) error {
	err := s.repo.ColumnRepo.Delete(deletingColumn)
	if err != nil {
		return err
	}

	return nil
}
