package handler

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/middleware"
	h_project "github.com/gznrf/go_task_tracker.back.git/pkg/handler/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
)

type Handler struct {
	authHandler       *h_auth.HAuth
	middleWareHandler *h_middleware.HMiddleWare
	userHandler       *h_user.HUser
	projectHandler    *h_project.HProject
	service           service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		authHandler:       h_auth.NewHAuth(services),
		middleWareHandler: h_middleware.NewMiddleWare(),
		userHandler:       h_user.NewHUser(services),
		projectHandler:    h_project.NewHProject(services),
		service:           *services,
	}
}

func (h *Handler) InitRoutes() *http.Handler {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.authHandler.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.authHandler.SignIn).Methods("POST")
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
		}
	}

	a := applyCORS(router)

	return &a
}

func applyCORS(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:63342",
			"http://127.0.0.1:63342",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(h)
}
