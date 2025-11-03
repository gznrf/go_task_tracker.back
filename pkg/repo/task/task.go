package r_task

import (
	"github.com/gznrf/go_task_tracker.back.git/models/task"
	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(data *m_task.CreateRequest) (*m_task.CreateResponse, error) {
	var output *m_task.CreateResponse
	output = new(m_task.CreateResponse)

	row := r.db.QueryRow(createQuery, data.ColumnId, data.CreatorId, data.ExecutorId, data.Number, data.Name, data.Description)
	if err := row.Scan(&output.CreatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *TaskRepo) Get() (*m_task.GetResponse, error) {
	var output *m_task.GetResponse
	output = new(m_task.GetResponse)

	err := r.db.Select(&output.TasksList, getQuery)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *TaskRepo) GetByColumnId(data *m_task.GetByColumnIdRequest) (*m_task.GetByColumnIdResponse, error) {
	var output *m_task.GetByColumnIdResponse
	output = new(m_task.GetByColumnIdResponse)

	err := r.db.Get(&output.TasksList, getByColumnIdQuery, data.ColumnId)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *TaskRepo) GetById(data *m_task.GetByIdRequest) (*m_task.GetByIdResponse, error) {
	var output *m_task.GetByIdResponse
	output = new(m_task.GetByIdResponse)

	err := r.db.Get(&output.Task, getByIdQuery, data.Id)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *TaskRepo) Update(data *m_task.UpdateRequest) (*m_task.UpdateResponse, error) {
	var output *m_task.UpdateResponse
	output = new(m_task.UpdateResponse)

	row := r.db.QueryRow(updateQuery, data.Id, data.ColumnId, data.CreatorId, data.ExecutorId, data.Number, data.Name, data.Description)
	if err := row.Scan(&output.UpdatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *TaskRepo) Delete(data *m_task.DeleteRequest) (*m_task.DeleteResponse, error) {
	var output *m_task.DeleteResponse
	output = new(m_task.DeleteResponse)

	row := r.db.QueryRow(deleteQuery, data.Id)
	if err := row.Scan(&output.DeletedId); err != nil {
		return nil, err
	}

	return output, nil
}
