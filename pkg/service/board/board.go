package s_board

import (
	"github.com/gznrf/go_task_tracker.back.git/models/board"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
)

type BoardService struct {
	repo *repo.Repo
}

func NewBoardService(repo *repo.Repo) *BoardService {
	return &BoardService{repo: repo}
}

func (s *BoardService) Create(creatingBoard *m_board.Board) (int64, error) {
	createdId, err := s.repo.BoardRepo.Create(creatingBoard)
	if err != nil {
		return 0, err
	}
	return createdId, nil
}
func (s *BoardService) Get() ([]m_board.Board, error) {
	outputBoards, err := s.repo.BoardRepo.Get()
	if err != nil {
		return nil, err
	}
	return outputBoards, nil
}
func (s *BoardService) GetById(boardId int64) (m_board.Board, error) {
	outputBoard, err := s.repo.BoardRepo.GetById(boardId)
	if err != nil {
		return m_board.Board{}, err
	}
	return outputBoard, nil
}
func (s *BoardService) GetByProjectId(projectId int64) ([]m_board.Board, error) {
	outputBoards, err := s.repo.BoardRepo.GetByProjectId(projectId)
	if err != nil {
		return nil, err
	}
	return outputBoards, nil
}
func (s *BoardService) Update(updatingBoard *m_board.Board) error {
	err := s.repo.BoardRepo.Update(updatingBoard)
	if err != nil {
		return err
	}
	return nil
}
func (s *BoardService) Delete(boardId int64) error {
	err := s.repo.BoardRepo.Delete(boardId)
	if err != nil {
		return err
	}
	return nil
}
