package r_projects

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/models/project"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo/utils"
	"github.com/jmoiron/sqlx"
)

type ProjectRepo struct {
	db *sqlx.DB
}

func NewProjectRepo(db *sqlx.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (r *ProjectRepo) Create(creatorId int64, creatingProject *m_project.Project) (int64, error) {
	var createdId int64
	query := fmt.Sprintf(`INSERT INTO %s (name, description, owner_id) VALUES ($1, $2, $3) RETURNING id`, app.ProjectTable)

	row := r.db.QueryRow(query, creatingProject.Name, creatingProject.Description, creatorId)
	if err := row.Scan(&createdId); err != nil {
		return 0, err
	}
	err := r_utils.CreateConnection(r.db, app.ProjectsUsers, createdId, creatorId)
	if err != nil {
		return 0, err
	}

	return createdId, nil
}
func (r *ProjectRepo) Get() ([]m_project.Project, error) {
	var outputProjects []m_project.Project
	query := fmt.Sprintf(`SELECT * FROM %s`, app.ProjectTable)
	err := r.db.Select(&outputProjects, query)
	if err != nil {
		return nil, err
	}

	return outputProjects, nil
}
func (r *ProjectRepo) GetByUserId(userId int64) ([]m_project.Project, error) {
	var outputProjects []m_project.Project

	query := fmt.Sprintf(`SELECT * FROM %s WHERE owner_id = $1`, app.ProjectTable)
	err := r.db.Get(&outputProjects, query, userId)
	if err != nil {
		return []m_project.Project{}, err
	}

	return outputProjects, nil
}
func (r *ProjectRepo) GetById(projectId int64) (m_project.Project, error) {
	var outputProject m_project.Project
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.ProjectTable)
	err := r.db.Get(&outputProject, query, projectId)
	if err != nil {
		return m_project.Project{}, err
	}

	return outputProject, nil
}
func (r *ProjectRepo) Update(updatingProject *m_project.Project) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1, descriprion = $2 WHERE id = $3`, app.ProjectTable)
	_, err := r.db.Exec(query, updatingProject.Name, updatingProject.Description, updatingProject.Id)

	return err
}
func (r *ProjectRepo) Delete(projectId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, app.ProjectTable)
	_, err := r.db.Exec(query, projectId)
	if err != nil {
		return err
	}

	err = r_utils.RemoveConnection(r.db, app.ProjectsUsers, projectId, "from_id")
	if err != nil {
		return err
	}

	return nil
}
func (r *ProjectRepo) CheckIsOwner(creatorId int64, projectId int64) (bool, error) {
	var outputProject m_project.Project
	query := fmt.Sprintf(`SELECT owner_id FROM %s WHERE id = $1`, app.ProjectTable)
	err := r.db.Get(&outputProject, query, projectId)
	if err != nil {
		return false, err
	}

	if outputProject.OwnerId != creatorId {
		return false, nil
	}

	return true, nil
}
