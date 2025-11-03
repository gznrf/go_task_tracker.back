package m_column

import "github.com/gznrf/go_task_tracker.back.git/models"

type Column struct {
	Id      int `json:"id" db:"id"`
	BoardId int `json:"board_id" db:"board_id"`

	Name string `json:"name" db:"name"`

	models.SoftDelete
}
