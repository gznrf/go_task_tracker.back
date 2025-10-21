package h_board

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/gznrf/go_task_tracker.back.git/utils"
)

type HBoard struct {
	service *service.Service
}

func NewHBoard(service *service.Service) *HBoard {
	return &HBoard{service: service}
}

func (h *HBoard) Create(w http.ResponseWriter, r *http.Request) {
	var input m_board.Board

	userId, err := utils.GetUserIdFromCtx(r)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if input.Name == "" || input.ProjectId <= 0 {
		utils.WriteError(w, 400, errors.New("name and project_id is required"))
		return
	}

	createdId, err := h.service.BoardService.Create(&input)
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
func (h *HBoard) Get(w http.ResponseWriter, r *http.Request) {
	boardsList, err := h.service.BoardService.Get()
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
		"userId":      userId,
		"boards_list": boardsList,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
func (h *HBoard) GetById(w http.ResponseWriter, r *http.Request) {
	var input m_board.GetByIdResponse

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
	outputBoard, err := h.service.BoardService.GetById(input.Id)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"userId": userId,
		"board":  outputBoard,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *HBoard) GetByProjectId(w http.ResponseWriter, r *http.Request) {
	var input m_board.GetByProjectIdResponse

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
	if input.Id <= 0 {
		utils.WriteError(w, 400, errors.New("id is required"))
		return
	}
	boardList, err := h.service.BoardService.GetByProjectId(input.Id)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"board_list": boardList,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}

}

func (h *HBoard) Update(w http.ResponseWriter, r *http.Request) {
	var input m_board.Board

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
	if input.Id == 0 || input.Name == "" {
		utils.WriteError(w, 400, errors.New("board_id is required, updating_data is required"))
		return
	}
	err := h.service.BoardService.Update(&input)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"updated_id": input.Id,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}

func (h *HBoard) Delete(w http.ResponseWriter, r *http.Request) {
	var input m_board.DeleteResponse

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, 500, errors.New("internal server error"))
		return
	}
	if input.Id <= 0 {
		utils.WriteError(w, 400, errors.New("id is required"))
		return
	}
	err := h.service.BoardService.Delete(input.Id)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}
	if err := utils.WriteJson(w, 200, map[string]interface{}{
		"deleted_id": input.Id,
	}); nil != err {
		utils.WriteError(w, 500, err)
		return
	}
}
