package h_user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type UserHandler struct {
	service *service.Service
}

func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	var input *m_user.GetRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.UserService.Get()
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"data": output,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input *m_user.GetByIdRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.UserService.GetById(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"data": output,
	}); err != nil {
		utils.WriteError(w, 500, err)
	}
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input *m_user.UpdateRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.UserService.Update(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"data": output,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input *m_user.DeleteRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.UserService.Delete(input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"data": output,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
