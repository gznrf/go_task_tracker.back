package h_auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type HAuth struct {
	service *service.Service
}

func NewHAuth(service *service.Service) *HAuth {
	return &HAuth{service: service}
}

func (h *HAuth) SignUp(w http.ResponseWriter, r *http.Request) {
	var input m_auth.User

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
	}

	if input.Username == "" || input.Password == "" {
		utils.WriteError(w, 400, errors.New("username or password is empty"))
	}

	createdId, err := h.service.AuthService.CreateUser(input.Username, input.Password)
	if err != nil {
		utils.WriteError(w, 500, err)
	}

	token, err := utils.GenerateToken(createdId)

	if err := utils.WriteJson(w, http.StatusOK, fmt.Sprintf("created id: %s \ntoken: %s", createdId, token)); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
	}
}

func (h *HAuth) SignIn(w http.ResponseWriter, r *http.Request) {
	var input m_auth.User

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
	}

	if input.Username == "" || input.Password == "" {
		utils.WriteError(w, 400, errors.New("username or password is empty"))
	}

	createdId, err := h.service.AuthService.LoginUser(input.Username, input.Password)
	if err != nil {
		utils.WriteError(w, 500, err)
	}

	token, err := utils.GenerateToken(createdId)
	if err != nil {
		utils.WriteError(w, 500, err)
	}

	if err := utils.WriteJson(w, http.StatusOK, fmt.Sprintf("Token: %s", token)); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
	}
}
