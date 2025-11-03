package m_board

import "github.com/gznrf/go_task_tracker.back.git/models"

type Board struct {
	Id        int64 `json:"id" db:"id"`
	ProjectId int64 `json:"project_id" db:"project_id"`

	Name string `json:"name" db:"name"`

	models.SoftDelete
}
