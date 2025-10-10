package handler

import (
	"github.com/gorilla/mux"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler/middleware"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
)

type Handler struct {
	authHandler       *h_auth.HAuth
	middleWareHandler *h_middleware.HMiddleWare
	service           service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		authHandler:       h_auth.NewHAuth(services),
		middleWareHandler: h_middleware.NewMiddleWare(),
		service:           *services,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.Use(h.middleWareHandler.UserIdentity)
		auth.HandleFunc("/sign-up", h.authHandler.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.authHandler.SignIn).Methods("POST")
	}

	api := router.PathPrefix("/api").Subrouter()
	{
		api.Use(h.middleWareHandler.UserIdentity)
	}

	return router
}
