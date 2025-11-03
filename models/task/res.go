package m_task

type CreateResponse struct {
	CreatedId int64 `json:"created_id"`
}

type GetResponse struct {
	TasksList []Task `json:"tasks_list"`
}

type GetByColumnIdResponse struct {
	TasksList []Task `json:"tasks_list" `
}

type GetByIdResponse struct {
	Task Task `json:"task" `
}

type UpdateResponse struct {
	UpdatedId int64 `json:"updated_id"`
}

type DeleteResponse struct {
	DeletedId int64 `json:"deleted_id"`
}
