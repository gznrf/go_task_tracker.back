package r_user

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Get() ([]m_user.User, error) {
	var outputUsers []m_user.User
	query := fmt.Sprintf(`SELECT * FROM %s`, app.UserTable)
	err := r.db.Select(&outputUsers, query)
	if err != nil {
		return nil, err
	}

	return outputUsers, nil
}
func (r *UserRepo) GetById(userID int64) (m_user.User, error) {
	var outputUser m_user.User
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.UserTable)
	err := r.db.Get(&outputUser, query, userID)
	if err != nil {
		return m_user.User{}, err
	}

	return outputUser, nil
}
func (r *UserRepo) Update(updatingData *m_user.UpdateUserRequest) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1 WHERE id = $2`, app.UserTable)
	_, err := r.db.Exec(query, updatingData.Username, updatingData.UserId)

	return err
}
func (r *UserRepo) Delete(userId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, app.UserTable)
	_, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}
