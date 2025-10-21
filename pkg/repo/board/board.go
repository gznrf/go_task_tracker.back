package r_board

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/jmoiron/sqlx"
)

type BoardRepo struct {
	db *sqlx.DB
}

func NewBoardRepo(db *sqlx.DB) *BoardRepo {
	return &BoardRepo{db: db}
}

func (r *BoardRepo) Create(creatingBoard *m_board.Board) (int64, error) {
	var createdId int64
	query := fmt.Sprintf(`INSERT INTO %s (name, project_id) VALUES ($1, $2) RETURNING id`, app.BoardTable)

	row := r.db.QueryRow(query, creatingBoard.Name, creatingBoard.ProjectId)
	if err := row.Scan(&createdId); err != nil {
		return 0, err
	}

	return createdId, nil
}
func (r *BoardRepo) Get() ([]m_board.Board, error) {
	var outputBoards []m_board.Board
	query := fmt.Sprintf(`SELECT * FROM %s`, app.BoardTable)
	err := r.db.Select(&outputBoards, query)
	if err != nil {
		return nil, err
	}

	return outputBoards, nil
}
func (r *BoardRepo) GetById(boardId int64) (m_board.Board, error) {
	var outputBoard m_board.Board

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.BoardTable)
	err := r.db.Get(&outputBoard, query, boardId)
	if err != nil {
		return m_board.Board{}, err
	}

	return outputBoard, nil

}
func (r *BoardRepo) GetByProjectId(projectId int64) ([]m_board.Board, error) {
	var outputBoards []m_board.Board

	query := fmt.Sprintf(`SELECT * FROM %s WHERE project_id = $1`, app.BoardTable)
	err := r.db.Select(&outputBoards, query, projectId)
	if err != nil {
		return nil, err
	}

	return outputBoards, nil
}
func (r *BoardRepo) Update(updatingBoard *m_board.Board) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1 WHERE id = $2`, app.BoardTable)
	_, err := r.db.Exec(query, updatingBoard.Name, updatingBoard.Id)

	return err
}
func (r *BoardRepo) Delete(boardId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, app.BoardTable)
	_, err := r.db.Exec(query, boardId)
	if err != nil {
		return err
	}

	return nil
}
