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
	var input *m_auth.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.AuthService.RegisterUser(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"data": output,
	}); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
}
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input *m_auth.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.AuthService.LoginUser(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	token, err := utils.GenerateToken(output.UserId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	utils.SetTokenToCookie(w, token)

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"data": output,
	}); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
}
