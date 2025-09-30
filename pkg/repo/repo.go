package repo

import "github.com/jmoiron/sqlx"

type Repo struct {
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{}
}
