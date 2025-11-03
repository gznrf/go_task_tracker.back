package h_board

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type BoardHandler struct {
	service *service.Service
}

func NewBoardHandler(service *service.Service) *BoardHandler {
	return &BoardHandler{service: service}
}

func (h *BoardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input *m_board.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.BoardService.Create(input)
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
func (h *BoardHandler) Get(w http.ResponseWriter, r *http.Request) {
	var input *m_board.GetRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	output, err := h.service.BoardService.Get()
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
func (h *BoardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var input *m_board.GetByIdRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.BoardService.GetById(input)
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
func (h *BoardHandler) GetByProjectId(w http.ResponseWriter, r *http.Request) {
	var input *m_board.GetByProjectIdRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.BoardService.GetByProjectId(input)
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
func (h *BoardHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input *m_board.UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.BoardService.Update(input)
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
func (h *BoardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input *m_board.DeleteRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}

	output, err := h.service.BoardService.Delete(input)
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
