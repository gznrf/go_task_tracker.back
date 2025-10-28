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
	var input m_column.CreateRequest

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	createdId, err := h.service.ColumnService.Create(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id":    userId,
		"created_id": createdId,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ColumnHandler) Get(w http.ResponseWriter, r *http.Request) {
	var _ m_column.GetRequest

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	columnsList, err := h.service.ColumnService.Get()
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id":      userId,
		"columns_list": columnsList,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ColumnHandler) GetByBoardId(w http.ResponseWriter, r *http.Request) {
	var input m_column.GetByBoardIdRequest

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	columnsList, err := h.service.ColumnService.GetByBoardId(input.BoardId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id":      userId,
		"board_id":     input.BoardId,
		"columns_list": columnsList,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ColumnHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input m_column.GetByIdRequest

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	outputColumn, err := h.service.ColumnService.GetById(input.ColumnId)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id": userId,
		"column":  outputColumn,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

}
func (h *ColumnHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input m_column.UpdateRequest

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = h.service.ColumnService.Update(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id":    userId,
		"updated_id": input.Id,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *ColumnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input m_column.DeleteRequest

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	err = h.service.ColumnService.Delete(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"user_id":    userId,
		"deleted_id": input.ColumnId,
	}); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
}
