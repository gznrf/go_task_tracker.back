package service

import (
	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/column"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/user"
)

type Auth interface {
	CreateUser(username string, password string) (int64, error)
	LoginUser(username string, password string) (int64, error)
}
type Middleware interface {
}

type User interface {
	Get() ([]m_user.User, error)
	GetById(userId int64) (m_user.User, error)
	Update(inputData *m_user.UpdateUserRequest) error
	Delete(userId int64) error
}

type Project interface {
	Create(creatorId int64, creatingProject *m_project.Project) (int64, error)
	Get() ([]m_project.Project, error)
	GetByUserId(userId int64) ([]m_project.Project, error)
	GetById(projectId int64) (m_project.Project, error)
	Update(creatorId int64, updatingProject *m_project.Project) error
	Delete(creatorId int64, projectId int64) error
	CheckIsOwner(creatorId int64, projectId int64) (bool, error)
}

type Board interface {
	Create(creatingBoard *m_board.Board) (int64, error)
	Get() ([]m_board.Board, error)
	GetById(boardId int64) (m_board.Board, error)
	GetByProjectId(projectId int64) ([]m_board.Board, error)
	Update(updatingBoard *m_board.Board) error
	Delete(boardId int64) error
} //короче тут три эмплементации

type Column interface {
	Create(creatingColumn *m_column.CreateRequest) (int64, error)
	Get() ([]m_column.Column, error)
	GetByBoardId(boardId int64) ([]m_column.Column, error)
	GetById(columnId int64) (m_column.GetByIdResponse, error)
	Update(updatingColumn *m_column.UpdateRequest) error
	Delete(deletingColumn *m_column.DeleteRequest) error
}

type Service struct {
	AuthService    Auth
	UserService    User
	ProjectService Project
	BoardService   Board
	ColumnService  Column
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		AuthService:    s_auth.NewAuthService(repo),
		UserService:    s_user.NewUserService(repo),
		ProjectService: s_project.NewProjectService(repo),
		BoardService:   s_board.NewBoardService(repo),
		ColumnService:  s_column.NewColumnService(repo),
	}
}
