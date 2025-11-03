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

func (s *BoardService) Create(data *m_board.CreateRequest) (*m_board.CreateResponse, error) {
	var output *m_board.CreateResponse

	output, err := s.repo.BoardRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *BoardService) Get() (*m_board.GetResponse, error) {
	var output *m_board.GetResponse

	output, err := s.repo.BoardRepo.Get()
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *BoardService) GetById(data *m_board.GetByIdRequest) (*m_board.GetByIdResponse, error) {
	var output *m_board.GetByIdResponse

	output, err := s.repo.BoardRepo.GetById(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *BoardService) GetByProjectId(data *m_board.GetByProjectIdRequest) (*m_board.GetByProjectIdResponse, error) {
	var output *m_board.GetByProjectIdResponse

	output, err := s.repo.BoardRepo.GetByProjectId(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *BoardService) Update(data *m_board.UpdateRequest) (*m_board.UpdateResponse, error) {
	var output *m_board.UpdateResponse

	output, err := s.repo.BoardRepo.Update(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *BoardService) Delete(data *m_board.DeleteRequest) (*m_board.DeleteResponse, error) {
	var output *m_board.DeleteResponse

	output, err := s.repo.BoardRepo.Delete(data)
	if err != nil {
		return nil, err
	}
	return output, nil
}
