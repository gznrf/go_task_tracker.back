package h_column

import (
	"encoding/json"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/column"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type ColumnHandler struct {
	service *service.Service
}

func NewColumnHandler(service *service.Service) *ColumnHandler {
	return &ColumnHandler{service: service}
}

func (h *ColumnHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input *m_column.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ColumnService.Create(input)
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
func (h *ColumnHandler) Get(w http.ResponseWriter, r *http.Request) {
	var input *m_column.GetRequest

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ColumnService.Get()
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
func (h *ColumnHandler) GetByBoardId(w http.ResponseWriter, r *http.Request) {
	var input *m_column.GetByBoardIdRequest

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ColumnService.GetByBoardId(input)
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
func (h *ColumnHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input *m_column.GetByIdRequest

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ColumnService.GetById(input)
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
func (h *ColumnHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input *m_column.UpdateRequest

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ColumnService.Update(input)
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
func (h *ColumnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input *m_column.DeleteRequest

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.ColumnService.Delete(input)
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
