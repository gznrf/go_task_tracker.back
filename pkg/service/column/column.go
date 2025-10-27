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
	return 0, nil
}
func (s *ColumnService) Get() ([]m_column.Column, error) {
	return nil, nil
}
func (s *ColumnService) GetById(columnId int64) (m_column.Column, error) {
	return m_column.Column{}, nil
}
func (s *ColumnService) GetByBoardId(boardId int64) ([]m_column.Column, error) {
	return nil, nil
}
func (s *ColumnService) Update(request *m_column.UpdateRequest) error {
	return nil
}
func (s *ColumnService) Delete(columnId int64) error {
	return nil
}
