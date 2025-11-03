package r_column

import (
	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/jmoiron/sqlx"
)

type ColumnRepo struct {
	db *sqlx.DB
}

func NewColumnRepo(db *sqlx.DB) *ColumnRepo {
	return &ColumnRepo{db: db}
}

func (r *ColumnRepo) Create(data *m_column.CreateRequest) (*m_column.CreateResponse, error) {
	var output *m_column.CreateResponse
	output = new(m_column.CreateResponse)

	row := r.db.QueryRow(createQuery, data.BoardId, data.Name)
	if err := row.Scan(&output.CreatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ColumnRepo) Get() (*m_column.GetResponse, error) {
	var output *m_column.GetResponse
	output = new(m_column.GetResponse)

	err := r.db.Select(&output.ColumnsList, getQuery)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ColumnRepo) GetByBoardId(data *m_column.GetByBoardIdRequest) (*m_column.GetByBoardIdResponse, error) {
	var output *m_column.GetByBoardIdResponse
	output = new(m_column.GetByBoardIdResponse)

	err := r.db.Select(&output.ColumnsList, getByBoardIdQuery, data.BoardId)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ColumnRepo) GetById(data *m_column.GetByIdRequest) (*m_column.GetByIdResponse, error) {
	var output *m_column.GetByIdResponse
	output = new(m_column.GetByIdResponse)

	err := r.db.Get(&output.Column, getByIdQuery, data.Id)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ColumnRepo) Update(data *m_column.UpdateRequest) (*m_column.UpdateResponse, error) {
	var output *m_column.UpdateResponse
	output = new(m_column.UpdateResponse)

	row := r.db.QueryRow(updateQuery, data.BoardId, data.Name, data.Id)
	if err := row.Scan(&output.UpdatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ColumnRepo) Delete(data *m_column.DeleteRequest) (*m_column.DeleteResponse, error) {
	var output *m_column.DeleteResponse
	output = new(m_column.DeleteResponse)

	row := r.db.QueryRow(deleteQuery, data.Id)
	if err := row.Scan(&output.DeletedId); err != nil {
		return nil, err
	}

	return output, nil
}
