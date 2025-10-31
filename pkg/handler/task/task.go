package h_task

import (
	"net/http"

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

	if err := utils.WriteJson(w, 200, map[string]interface{}{}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {

	if err := utils.WriteJson(w, 200, map[string]interface{}{}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *TaskHandler) GetById(w http.ResponseWriter, r *http.Request) {

	if err := utils.WriteJson(w, 200, map[string]interface{}{}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *TaskHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {

	if err := utils.WriteJson(w, 200, map[string]interface{}{}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {

	if err := utils.WriteJson(w, 200, map[string]interface{}{}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {

	if err := utils.WriteJson(w, 200, map[string]interface{}{}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
