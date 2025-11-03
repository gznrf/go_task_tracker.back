package m_task

import (
	"github.com/gznrf/go_task_tracker.back.git/models"
)

type Task struct {
	Id         int64 `json:"id" db:"id"`
	ColumnId   int64 `json:"column_id" db:"column_id"`
	CreatorId  int64 `json:"creator_id" db:"creator_id"`
	ExecutorId int64 `json:"executor_id" db:"executor_id"`

	Number      int64  `json:"number" db:"number"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	models.SoftDelete
}
