package repo

import (
	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/column"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/project"
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

type Board interface {
	Create(creatingBoard *m_board.Board) (int64, error)
	Get() ([]m_board.Board, error)
	GetById(boardId int64) (m_board.Board, error)
	GetByProjectId(projectId int64) ([]m_board.Board, error)
	Update(updatingBoard *m_board.Board) error
	Delete(boardId int64) error
}

type Column interface {
	Create(creatingColumn *m_column.CreateRequest) (int64, error)
	Get() ([]m_column.Column, error)
	GetByBoardId(boardId int64) ([]m_column.Column, error)
	GetById(columnId int64) (m_column.GetByIdResponse, error)
	Update(request *m_column.UpdateRequest) error
	Delete(columnId int64) error
}

type Repo struct {
	AuthRepo    Auth
	UserRepo    User
	ProjectRepo Project
	BoardRepo   Board
	ColumnRepo  Column
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		AuthRepo:    r_auth.NewAuthRepo(db),
		UserRepo:    r_user.NewUserRepo(db),
		ProjectRepo: r_projects.NewProjectRepo(db),
		BoardRepo:   r_board.NewBoardRepo(db),
		ColumnRepo:  r_column.NewColumnRepo(db),
	}
}
