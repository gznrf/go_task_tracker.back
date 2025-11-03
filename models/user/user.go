package m_user

import "github.com/gznrf/go_task_tracker.back.git/models"

type User struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`

	models.SoftDelete
}
