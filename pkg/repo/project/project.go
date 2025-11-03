package r_projects

import (
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/jmoiron/sqlx"
)

type ProjectRepo struct {
	db *sqlx.DB
}

func NewProjectRepo(db *sqlx.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (r *ProjectRepo) Create(data *m_project.CreateRequest) (*m_project.CreateResponse, error) {
	var output *m_project.CreateResponse
	output = new(m_project.CreateResponse)

	row := r.db.QueryRow(createQuery, data.OwnerId, data.Name, data.Description)
	if err := row.Scan(&output.CreatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ProjectRepo) Get() (*m_project.GetResponse, error) {
	var output *m_project.GetResponse
	output = new(m_project.GetResponse)

	err := r.db.Select(&output.ProjectsList, getQuery)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ProjectRepo) GetByUserId(data *m_project.GetByUserIdRequest) (*m_project.GetByUserIdResponse, error) {
	var output *m_project.GetByUserIdResponse
	output = new(m_project.GetByUserIdResponse)

	err := r.db.Select(&output.ProjectsList, getByUserIdQuery, data.UserId)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ProjectRepo) GetById(data *m_project.GetByIdRequest) (*m_project.GetByIdResponse, error) {
	var output *m_project.GetByIdResponse
	output = new(m_project.GetByIdResponse)

	err := r.db.Get(&output, getByIdQuery, data.Id)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ProjectRepo) Update(data *m_project.UpdateRequest) (*m_project.UpdateResponse, error) {
	var output *m_project.UpdateResponse
	output = new(m_project.UpdateResponse)

	row := r.db.QueryRow(updateQuery, data.Id, data.OwnerId, data.Name, data.Description)
	if err := row.Scan(&output.UpdatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *ProjectRepo) Delete(data *m_project.DeleteRequest) (*m_project.DeleteResponse, error) {
	var output *m_project.DeleteResponse
	output = new(m_project.DeleteResponse)

	row := r.db.QueryRow(deleteQuery, data.Id)
	if err := row.Scan(&output.DeletedId); err != nil {
		return nil, err
	}

	return output, nil
}
