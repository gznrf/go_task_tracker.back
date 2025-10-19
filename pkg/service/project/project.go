package s_project

import (
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
)

type ProjectService struct {
	repo *repo.Repo
}

func NewProjectService(repo *repo.Repo) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) Create(creatorId int64, creatingProject *m_project.Project) (int64, error) {
	createdId, err := s.repo.ProjectRepo.Create(creatorId, creatingProject)
	if err != nil {
		return 0, err
	}

	return createdId, nil
}
func (s *ProjectService) Get() ([]m_project.Project, error) {
	return s.repo.ProjectRepo.Get()
}
func (s *ProjectService) GetByUserId(userId int64) ([]m_project.Project, error) {
	usersProjects, err := s.repo.ProjectRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return usersProjects, nil
}
func (s *ProjectService) GetById(projectId int64) (m_project.Project, error) {
	outputProject, err := s.repo.ProjectRepo.GetById(projectId)
	if err != nil {
		return m_project.Project{}, nil
	}
	return outputProject, nil
}
func (s *ProjectService) Update(creatorId int64, updatingProject *m_project.Project) error {
	isOwner, err := s.CheckIsOwner(creatorId, updatingProject.Id)
	if err != nil {
		return err
	}
	// if owner --> update project else --> doesn't update
	if isOwner {
		err = s.repo.ProjectRepo.Update(updatingProject)
		if err != nil {
			return err
		}
	}

	return nil
}
func (s *ProjectService) Delete(creatorId int64, projectId int64) error {
	isOwner, err := s.CheckIsOwner(creatorId, projectId)
	if err != nil {
		return err
	}
	// if owner --> delete project else --> doesn't delete
	if isOwner {
		err = s.repo.ProjectRepo.Delete(projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ProjectService) CheckIsOwner(creatorId int64, projectId int64) (bool, error) {
	isOwner, err := s.repo.ProjectRepo.CheckIsOwner(creatorId, projectId)
	if err != nil {
		return false, err
	}
	if isOwner {
		return true, nil
	}
	return false, nil
}
