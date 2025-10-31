package m_task

import "time"

type Task struct {
	Id         int64 `json:"id" db:"id"`
	ColumnId   int64 `json:"column_id" db:"column_id"`
	CreatorId  int64 `json:"creator_id" db:"creator_id"`
	ExecutorId int64 `json:"executor_id" db:"executor_id"`

	Number      int64  `json:"number" db:"number"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
