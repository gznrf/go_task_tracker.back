package r_column

import (
	m_column "github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/jmoiron/sqlx"
)

type ColumnRepo struct {
	db *sqlx.DB
}

func NewColumnRepo(db *sqlx.DB) *ColumnRepo {
	return &ColumnRepo{db: db}
}

func (r *ColumnRepo) Create(creatingColumn *m_column.CreateRequest) (int64, error) {
	return 0, nil
}
func (r *ColumnRepo) Get() ([]m_column.Column, error) {
	return nil, nil
}
func (r *ColumnRepo) GetByBoardId(boardId int64) ([]m_column.Column, error) {
	return nil, nil
}
func (r *ColumnRepo) GetById(columnId int64) (m_column.GetByIdResponse, error) {
	return m_column.GetByIdResponse{}, nil
}
func (r *ColumnRepo) Update(request *m_column.UpdateRequest) error {
	return nil
}
func (r *ColumnRepo) Delete(columnId int64) error {
	return nil
}
