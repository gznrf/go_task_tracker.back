package h_user

import (
	"encoding/json"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/user"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type HUser struct {
	service *service.Service
}

func NewHUser(service *service.Service) *HUser {
	return &HUser{service: service}
}

func (h *HUser) Get(w http.ResponseWriter, r *http.Request) {
	outputUsers, err := h.service.UserService.Get()
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"users": outputUsers,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *HUser) GetById(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	user, err := h.service.UserService.GetById(userId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	res := m_user.GetUserByIdResponse{
		User: user,
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user": res.User,
	}); err != nil {
		utils.WriteError(w, 500, err)
	}
}

func (h *HUser) Update(w http.ResponseWriter, r *http.Request) {
	var input m_user.UpdateUserRequest
	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	input.UserId = userId
	err = h.service.UserService.Update(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

}

func (h *HUser) Delete(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = h.service.UserService.Delete(userId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"is_user_deleted": "true",
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
