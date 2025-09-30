package handler

import (
	"github.com/gorilla/mux"
	handler "github.com/gznrf/go_task_tracker.back.git/pkg/handler/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
)

type Handler struct {
	auth     *handler.HAuth
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{

		auth.HandleFunc("/sign-up", h.auth.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.auth.SignIn).Methods("POST")
	}

	/*api := router.PathPrefix("/api").Subrouter()
	{
		//api.Use(h.userIdentity)
	}*/

	return router
}
