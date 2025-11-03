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
	var input *m_project.CreateRequest
	input = new(m_project.CreateRequest)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	userID, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	input.OwnerId = userID

	output, err := h.service.ProjectService.Create(input)
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
func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	var input *m_project.GetRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ProjectService.Get()
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
func (h *ProjectHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input *m_project.GetByIdRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.ProjectService.GetById(input)
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
func (h *ProjectHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	var input *m_project.GetByUserIdRequest
	input = new(m_project.GetByUserIdRequest)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	input.UserId = userId

	output, err := h.service.ProjectService.GetByUserId(input)
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
func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input *m_project.UpdateRequest
	input = new(m_project.UpdateRequest)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.ProjectService.Update(input)
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
func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input *m_project.DeleteRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.ProjectService.Delete(input)
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
