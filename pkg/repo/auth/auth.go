package r_auth

import (
	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) RegisterUser(data *m_auth.RegisterRequest) (*m_auth.RegisterResponse, error) {
	var output *m_auth.RegisterResponse
	output = new(m_auth.RegisterResponse)

	row := r.db.QueryRow(registerQuery, data.Name, data.Email, data.Password)
	if err := row.Scan(&output.CreatedId); err != nil {
		return nil, err
	}

	return output, nil
}
func (r *AuthRepo) LoginUser(data *m_auth.LoginRequest) (*m_auth.LoginResponse, string, error) {
	var output *m_auth.LoginResponse
	output = new(m_auth.LoginResponse)
	var outputUser m_auth.User

	err := r.db.Get(&outputUser, loginQuery, data.Email, data.Password)
	if err != nil {
		return nil, "", err
	}
	output.UserId = outputUser.Id

	return output, outputUser.Password, nil
}
