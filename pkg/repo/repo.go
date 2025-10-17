package repo

import (
	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/user"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(username string, password string) (createdId int64, err error)
	GetUser(username, hashedPassword string) (user m_auth.User, err error)
}

type User interface {
	Get() ([]m_user.User, error)
	GetById(userID int64) (m_user.User, error)
	Update(updatingData *m_user.UpdateUserRequest) error
	Delete(userId int64) error
}

type Repo struct {
	db       *sqlx.DB
	AuthRepo Auth
	UserRepo User
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		db:       db,
		AuthRepo: r_auth.NewAuthRepo(db),
		UserRepo: r_user.NewUserRepo(db),
	}
}
