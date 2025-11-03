package service

import (
	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/models/task"
	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/column"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/task"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service/user"
)

type Auth interface {
	RegisterUser(data *m_auth.RegisterRequest) (*m_auth.RegisterResponse, error)
	LoginUser(data *m_auth.LoginRequest) (*m_auth.LoginResponse, error)
}
type Middleware interface {
}
type User interface {
	Get() (*m_user.GetResponse, error)
	GetById(data *m_user.GetByIdRequest) (*m_user.GetByIdResponse, error)
	Update(data *m_user.UpdateRequest) (*m_user.UpdateResponse, error)
	Delete(data *m_user.DeleteRequest) (*m_user.DeleteResponse, error)
}
type Project interface {
	Create(data *m_project.CreateRequest) (*m_project.CreateResponse, error)
	Get() (*m_project.GetResponse, error)
	GetByUserId(data *m_project.GetByUserIdRequest) (*m_project.GetByUserIdResponse, error)
	GetById(data *m_project.GetByIdRequest) (*m_project.GetByIdResponse, error)
	Update(data *m_project.UpdateRequest) (*m_project.UpdateResponse, error)
	Delete(data *m_project.DeleteRequest) (*m_project.DeleteResponse, error)
}
type Board interface {
	Create(data *m_board.CreateRequest) (*m_board.CreateResponse, error)
	Get() (*m_board.GetResponse, error)
	GetById(data *m_board.GetByIdRequest) (*m_board.GetByIdResponse, error)
	GetByProjectId(data *m_board.GetByProjectIdRequest) (*m_board.GetByProjectIdResponse, error)
	Update(data *m_board.UpdateRequest) (*m_board.UpdateResponse, error)
	Delete(data *m_board.DeleteRequest) (*m_board.DeleteResponse, error)
}
type Column interface {
	Create(data *m_column.CreateRequest) (*m_column.CreateResponse, error)
	Get() (*m_column.GetResponse, error)
	GetByBoardId(data *m_column.GetByBoardIdRequest) (*m_column.GetByBoardIdResponse, error)
	GetById(data *m_column.GetByIdRequest) (*m_column.GetByIdResponse, error)
	Update(data *m_column.UpdateRequest) (*m_column.UpdateResponse, error)
	Delete(data *m_column.DeleteRequest) (*m_column.DeleteResponse, error)
}
type Task interface {
	Create(data *m_task.CreateRequest) (*m_task.CreateResponse, error)
	Get() (*m_task.GetResponse, error)
	GetByColumnId(data *m_task.GetByColumnIdRequest) (*m_task.GetByColumnIdResponse, error)
	GetById(data *m_task.GetByIdRequest) (*m_task.GetByIdResponse, error)
	Update(data *m_task.UpdateRequest) (*m_task.UpdateResponse, error)
	Delete(data *m_task.DeleteRequest) (*m_task.DeleteResponse, error)
}

type Service struct {
	AuthService    Auth
	UserService    User
	ProjectService Project
	BoardService   Board
	ColumnService  Column
	TaskService    Task
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		AuthService:    s_auth.NewAuthService(repo),
		UserService:    s_user.NewUserService(repo),
		ProjectService: s_project.NewProjectService(repo),
		BoardService:   s_board.NewBoardService(repo),
		ColumnService:  s_column.NewColumnService(repo),
		TaskService:    s_task.NewTaskService(repo),
	}
}
