package h_auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/auth"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type AuthHandler struct {
	service *service.Service
}

func NewAuthHandler(service *service.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input m_auth.RegisterRequest

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	if input.Username == "" || input.Password == "" {
		utils.WriteError(w, 400, errors.New("username or password is empty"))
		return
	}

	createdId, err := h.service.AuthService.CreateUser(input.Username, input.Password)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	token, err := utils.GenerateToken(createdId)

	res := m_auth.RegisterResponse{
		Token:         token,
		CreatedUserId: createdId,
	}
	if err := utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"created_id": res.CreatedUserId,
		"token":      res.Token,
	}); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input m_auth.LoginRequest

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	if input.Username == "" || input.Password == "" {
		utils.WriteError(w, 400, errors.New("username or password is empty"))
		return
	}

	createdId, err := h.service.AuthService.LoginUser(input.Username, input.Password)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	token, err := utils.GenerateToken(createdId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	res := m_auth.LoginResponse{
		Token: token,
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"token": res.Token,
	}); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
}
