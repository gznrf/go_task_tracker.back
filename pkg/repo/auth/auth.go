package r_auth

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(username string, hashedPassword string) (int64, error) {
	var createdId int64
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash) VALUES ($1, $2) RETURNING id", app.UserTable)

	row := r.db.QueryRow(query, username, hashedPassword)
	if err := row.Scan(&createdId); err != nil {
		return 0, err
	}

	return createdId, nil
}

func (r *AuthRepo) GetUser(username, hashedPassword string) (m_auth.User, error) {
	var outputPass m_auth.ChangePasswordResponse
	var outputUser m_auth.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 AND password_hash = $2", app.UserTable)
	err := r.db.Get(&outputPass, query, username, hashedPassword)
	outputUser = m_auth.User{
		Username: username,
		Password: outputPass.PasswordHash,
	}
	return outputUser, err
}
