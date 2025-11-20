package handler

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/column"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/middleware"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/task"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
)

type Handler struct {
	authHandler       *h_auth.AuthHandler
	middleWareHandler *h_middleware.MiddleWareHandler
	userHandler       *h_user.UserHandler
	projectHandler    *h_project.ProjectHandler
	boardHandler      *h_board.BoardHandler
	columnHandler     *h_column.ColumnHandler
	taskHandler       *h_task.TaskHandler
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		authHandler:       h_auth.NewAuthHandler(services),
		middleWareHandler: h_middleware.NewMiddleWareHandler(),
		userHandler:       h_user.NewUserHandler(services),
		projectHandler:    h_project.NewProjectHandler(services),
		boardHandler:      h_board.NewBoardHandler(services),
		columnHandler:     h_column.NewColumnHandler(services),
		taskHandler:       h_task.NewTaskHandler(services),
	}
}

func (h *Handler) InitRoutes() *http.Handler {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.authHandler.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.authHandler.SignIn).Methods("POST")
	}
	auth2 := router.PathPrefix("/auth2").Subrouter()
	auth2.Use(h.middleWareHandler.UserIdentity)
	{
		auth2.HandleFunc("/isAuth", h.authHandler.IsAuth).Methods("POST")
		auth2.HandleFunc("/logout", h.authHandler.Logout).Methods("POST")
	}

	api := router.PathPrefix("/api").Subrouter()
	{
		api.Use(h.middleWareHandler.UserIdentity)

		user := api.PathPrefix("/user").Subrouter()
		{
			user.HandleFunc("/get", h.userHandler.Get).Methods("GET")
			user.HandleFunc("/getById", h.userHandler.GetById).Methods("GET")
			user.HandleFunc("/update", h.userHandler.Update).Methods("PATCH")
			user.HandleFunc("/delete", h.userHandler.Delete).Methods("DELETE")
		}
		project := api.PathPrefix("/project").Subrouter()
		{
			project.HandleFunc("/create", h.projectHandler.Create).Methods("POST")
			project.HandleFunc("/get", h.projectHandler.Get).Methods("GET")
			project.HandleFunc("/getByUserId", h.projectHandler.GetByUserId).Methods("GET")
			project.HandleFunc("/getById", h.projectHandler.GetById).Methods("GET")
			project.HandleFunc("/update", h.projectHandler.Update).Methods("PATCH")
			project.HandleFunc("/delete", h.projectHandler.Delete).Methods("DELETE")

			board := project.PathPrefix("/board").Subrouter()
			{
				board.HandleFunc("/create", h.boardHandler.Create).Methods("POST")
				board.HandleFunc("/get", h.boardHandler.Get).Methods("GET")
				board.HandleFunc("/getByProjectId", h.boardHandler.GetByProjectId).Methods("GET")
				board.HandleFunc("/getById", h.boardHandler.GetById).Methods("GET")
				board.HandleFunc("/update", h.boardHandler.Update).Methods("PATCH")
				board.HandleFunc("/delete", h.boardHandler.Delete).Methods("DELETE")

				column := board.PathPrefix("/column").Subrouter()
				{
					column.HandleFunc("/create", h.columnHandler.Create).Methods("POST")
					column.HandleFunc("/get", h.columnHandler.Get).Methods("GET")
					column.HandleFunc("/getByBoardId", h.columnHandler.GetByBoardId).Methods("GET")
					column.HandleFunc("/getById", h.columnHandler.GetById).Methods("GET")
					column.HandleFunc("/update", h.columnHandler.Update).Methods("PATCH")
					column.HandleFunc("/delete", h.columnHandler.Delete).Methods("DELETE")

					task := column.PathPrefix("/task").Subrouter()
					{
						task.HandleFunc("/create", h.taskHandler.Create).Methods("POST")
						task.HandleFunc("/get", h.taskHandler.Get).Methods("GET")
						task.HandleFunc("/getByColumnId", h.taskHandler.GetByColumnId).Methods("GET")
						task.HandleFunc("/getById", h.taskHandler.GetById).Methods("GET")
						task.HandleFunc("/update", h.taskHandler.Update).Methods("PATCH")
						task.HandleFunc("/delete", h.taskHandler.Delete).Methods("DELETE")
					}

				}
			}
		}
	}

	a := applyCORS(router)

	return &a
}

func applyCORS(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://195.133.77.31:3321/",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(h)
	//
}
