package h_task

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/task"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type TaskHandler struct {
	service *service.Service
}

func NewTaskHandler(service *service.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input *m_task.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.TaskService.Create(input)
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
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	var input *m_task.GetRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.TaskService.Get()
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
func (h *TaskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input *m_task.GetByIdRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.TaskService.GetById(input)
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

func (h *TaskHandler) GetByColumnId(w http.ResponseWriter, r *http.Request) {
	var input *m_task.GetByColumnIdRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.TaskService.GetByColumnId(input)
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
func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input *m_task.UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.TaskService.Update(input)
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
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input *m_task.DeleteRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.TaskService.Delete(input)
	if err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"data": output,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
