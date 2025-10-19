package repo

import (
	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	m_project "github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/auth"
	r_projects "github.com/gznrf/go_task_tracker.back.git/pkg/repo/project"
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

type Project interface {
	Create(creatorId int64, creatingProject *m_project.Project) (int64, error)
	Get() ([]m_project.Project, error)
	GetByUserId(userId int64) ([]m_project.Project, error)
	GetById(projectId int64) (m_project.Project, error)
	Update(updatingProject *m_project.Project) error
	Delete(projectId int64) error
	CheckIsOwner(creatorId int64, projectId int64) (bool, error)
}

type Repo struct {
	db          *sqlx.DB
	AuthRepo    Auth
	UserRepo    User
	ProjectRepo Project
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		db:          db,
		AuthRepo:    r_auth.NewAuthRepo(db),
		UserRepo:    r_user.NewUserRepo(db),
		ProjectRepo: r_projects.NewProjectRepo(db),
	}
}
