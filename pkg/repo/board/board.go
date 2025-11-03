package r_board

import (
	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/jmoiron/sqlx"
)

type BoardRepo struct {
	db *sqlx.DB
}

func NewBoardRepo(db *sqlx.DB) *BoardRepo {
	return &BoardRepo{db: db}
}

func (r *BoardRepo) Create(data *m_board.CreateRequest) (*m_board.CreateResponse, error) {
	var output *m_board.CreateResponse
	output = new(m_board.CreateResponse)

	row := r.db.QueryRow(createQuery, data.ProjectId, data.Name)
	if err := row.Scan(&output.CreatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *BoardRepo) Get() (*m_board.GetResponse, error) {
	var output *m_board.GetResponse
	output = new(m_board.GetResponse)

	err := r.db.Select(&output.BoardsList, getQuery)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *BoardRepo) GetById(data *m_board.GetByIdRequest) (*m_board.GetByIdResponse, error) {
	var output *m_board.GetByIdResponse
	output = new(m_board.GetByIdResponse)

	err := r.db.Get(&output.Board, getByIdQuery, data.Id)
	if err != nil {
		return nil, err
	}

	return output, nil

}
func (r *BoardRepo) GetByProjectId(data *m_board.GetByProjectIdRequest) (*m_board.GetByProjectIdResponse, error) {
	var output *m_board.GetByProjectIdResponse
	output = new(m_board.GetByProjectIdResponse)

	err := r.db.Select(&output.BoardsList, getByProjectIdQuery, data.ProjectId)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *BoardRepo) Update(data *m_board.UpdateRequest) (*m_board.UpdateResponse, error) {
	var output *m_board.UpdateResponse
	output = new(m_board.UpdateResponse)

	row := r.db.QueryRow(updateQuery, data.ProjectId, data.Name, data.Id)
	if err := row.Scan(&output.UpdatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *BoardRepo) Delete(data *m_board.DeleteRequest) (*m_board.DeleteResponse, error) {
	var output *m_board.DeleteResponse
	output = new(m_board.DeleteResponse)

	row := r.db.QueryRow(deleteQuery, data.Id)
	if err := row.Scan(&output.DeletedId); err != nil {
		return nil, err
	}

	return output, nil
}
