package h_project

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type ProjectHandler struct {
	service *service.Service
}

func NewProjectHandler(service *service.Service) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input m_project.Project

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	input.OwnerId = userId

	if input.Name == "" || input.OwnerId <= 0 {
		utils.WriteError(w, 400, errors.New("name is required"))
		return
	}
	createdId, err := h.service.ProjectService.Create(userId, &input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id":    userId,
		"created_id": createdId,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	projectsList, err := h.service.ProjectService.Get()
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId":        userId,
		"projects_list": projectsList,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ProjectHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input m_project.GetByIdResponse

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
	if input.Id <= 0 {
		utils.WriteError(w, 400, errors.New("id is required"))
		return
	}
	outputProject, err := h.service.ProjectService.GetById(input.Id)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId":  userId,
		"project": outputProject,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

// GetByUserId id gets from token
func (h *ProjectHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	projectsList, err := h.service.ProjectService.GetByUserId(userId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId":        userId,
		"projects_list": projectsList,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input m_project.Project

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
	if input.Id == 0 || input.Name == "" || input.Description == "" {
		utils.WriteError(w, 400, errors.New("project_id is required, updating_data is required"))
		return
	}
	err = h.service.ProjectService.Update(userId, &input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId":     userId,
		"updated_id": input.Id,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input m_project.DeleteResponse

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
	if input.Id == 0 {
		utils.WriteError(w, 400, errors.New("id is required"))
		return
	}
	err = h.service.ProjectService.Delete(userId, input.Id)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId":     userId,
		"deleted_id": input.Id,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
