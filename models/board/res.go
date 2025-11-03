package m_board

type CreateResponse struct {
	CreatedId int64 `json:"created_id"`
}

type GetResponse struct {
	BoardsList []Board `json:"boards_list"`
}

type GetByProjectIdResponse struct {
	BoardsList []Board `json:"boards_list"`
}

type GetByIdResponse struct {
	Board Board `json:"board"`
}

type UpdateResponse struct {
	UpdatedId int64 `json:"updated_id"`
}

type DeleteResponse struct {
	DeletedId int64 `json:"deleted_id"`
}
