package repo

import (
	m_auth "github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/auth"
	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	CreateUser(username string, password string) (createdId int64, err error)
	GetUser(username, hashedPassword string) (user *m_auth.User, err error)
}

type Repo struct {
	db       *sqlx.DB
	AuthRepo *r_auth.AuthRepo
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		db:       db,
		AuthRepo: r_auth.NewAuthRepo(db),
	}
}
