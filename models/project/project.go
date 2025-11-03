package m_project

import "github.com/gznrf/go_task_tracker.back.git/models"

type Project struct {
	Id      int64 `json:"id" db:"id"`
	OwnerId int64 `json:"owner_id" db:"owner_id"`

	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	models.SoftDelete
}
