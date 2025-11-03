package m_board

type CreateRequest struct {
	ProjectId int64 `json:"project_id" db:"project_id"`

	Name string `json:"name" db:"name"`
}

type GetRequest struct {
}

type GetByProjectIdRequest struct {
	ProjectId int64 `json:"project_id" db:"project_id"`
}

type GetByIdRequest struct {
	Id int64 `json:"id" db:"id"`
}

type UpdateRequest struct {
	Id        int64 `json:"id" db:"id"`
	ProjectId int64 `json:"project_id" db:"project_id"`

	Name string `json:"name" db:"name"`
}

type DeleteRequest struct {
	Id int64 `json:"id" db:"id"`
}
