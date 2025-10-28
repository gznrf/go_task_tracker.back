package r_column

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/jmoiron/sqlx"
)

type ColumnRepo struct {
	db *sqlx.DB
}

func NewColumnRepo(db *sqlx.DB) *ColumnRepo {
	return &ColumnRepo{db: db}
}

func (r *ColumnRepo) Create(creatingColumn *m_column.CreateRequest) (int64, error) {
	var createdId int64
	query := fmt.Sprintf(`INSERT INTO %s (name, board_id) VALUES ($1, $2) RETURNING id`, app.ColumnTable)

	row := r.db.QueryRow(query, creatingColumn.Name, creatingColumn.BoardId)
	if err := row.Scan(&createdId); err != nil {
		return 0, err
	}

	return createdId, nil
}
func (r *ColumnRepo) Get() ([]m_column.Column, error) {
	var outputColumns []m_column.Column
	query := fmt.Sprintf(`SELECT * FROM %s`, app.ColumnTable)
	err := r.db.Select(&outputColumns, query)
	if err != nil {
		return nil, err
	}

	return outputColumns, nil
}
func (r *ColumnRepo) GetByBoardId(boardId int64) ([]m_column.Column, error) {
	var outputColumns []m_column.Column

	query := fmt.Sprintf(`SELECT * FROM %s WHERE board_id = $1`, app.ColumnTable)
	err := r.db.Select(&outputColumns, query, boardId)
	if err != nil {
		return nil, err
	}

	return outputColumns, nil
}
func (r *ColumnRepo) GetById(columnId int64) (m_column.GetByIdResponse, error) {
	var outputColumn m_column.GetByIdResponse
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.ColumnTable)
	err := r.db.Get(&outputColumn, query, columnId)
	if err != nil {
		return m_column.GetByIdResponse{}, err
	}

	return outputColumn, nil
}
func (r *ColumnRepo) Update(updatingColumn *m_column.UpdateRequest) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1, board_id = $2 WHERE id = $3`, app.ColumnTable)
	_, err := r.db.Exec(query, updatingColumn.Name, updatingColumn.BoardId, updatingColumn.Id)

	return err
}
func (r *ColumnRepo) Delete(deletingColumn *m_column.DeleteRequest) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, app.ColumnTable)
	_, err := r.db.Exec(query, deletingColumn.ColumnId)
	if err != nil {
		return err
	}

	return nil
}
