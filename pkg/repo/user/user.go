package r_user

import (
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Get() (*m_user.GetResponse, error) {
	var output *m_user.GetResponse
	output = new(m_user.GetResponse)

	err := r.db.Select(&output.UsersList, getQuery)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *UserRepo) GetById(data *m_user.GetByIdRequest) (*m_user.GetByIdResponse, error) {
	var output *m_user.GetByIdResponse
	output = new(m_user.GetByIdResponse)

	err := r.db.Get(&output.User, getByIdQuery, data.Id)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (r *UserRepo) Update(data *m_user.UpdateRequest) (*m_user.UpdateResponse, error) {
	var output *m_user.UpdateResponse
	output = new(m_user.UpdateResponse)

	row := r.db.QueryRow(updateQuery, data.Name, data.Email, data.Password, data.Id)
	if err := row.Scan(&output.UpdatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *UserRepo) Delete(data *m_user.DeleteRequest) (*m_user.DeleteResponse, error) {
	var output *m_user.DeleteResponse
	output = new(m_user.DeleteResponse)

	row := r.db.QueryRow(deleteQuery, data.Id)
	if err := row.Scan(&output.DeletedId); err != nil {
		return nil, err
	}

	return output, nil
}
