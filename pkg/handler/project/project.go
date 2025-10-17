package h_project

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type HProject struct {
	service *service.Service
}

func NewHProject(service *service.Service) *HProject {
	return &HProject{service: service}
}

func (h *HProject) Create(w http.ResponseWriter, r *http.Request) {
	var input m_project.Project
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if input.Name == "" || input.OwnerId == 0 {
		utils.WriteError(w, 400, errors.New("name and owner_id is required"))
		return
	}
	//TODO service.Create --> created_id
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		//TODO created_id
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *HProject) Get(w http.ResponseWriter, r *http.Request) {
	//TODO service.Get --> project_list

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId": userId,
		//TODO project_list
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *HProject) GetById(w http.ResponseWriter, r *http.Request) {
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
	if input.Id == 0 {
		utils.WriteError(w, 400, errors.New("id is required"))
		return
	}
	//TODO service.GetById --> project
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId": userId,
		//TODO project
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

// GetByUserId id gets from token
func (h *HProject) GetByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	//TODO service.GetByUserId --> services_list
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId": userId,
		//TODO services_list
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}

}
func (h *HProject) Update(w http.ResponseWriter, r *http.Request) {
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
	if input.Id == 0 || input.Name != "" || input.Description == "" {
		utils.WriteError(w, 400, errors.New("project_id is required, updating_data is required"))
		return
	}
	//TODO service.Update --> updated_id
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId": userId,
		//TODO updated_id
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *HProject) Delete(w http.ResponseWriter, r *http.Request) {
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
	//TODO service.Delete --> deleted_id
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId": userId,
		//TODO deleted_id
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
