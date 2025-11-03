package m_project

type CreateRequest struct {
	OwnerId int64 `json:"owner_id" db:"owner_id"` // takes from token

	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type GetRequest struct {
}

type GetByUserIdRequest struct {
	UserId int64 `json:"user_id"`
}

type GetByIdRequest struct {
	Id int64 `json:"id" db:"id"`
}

type UpdateRequest struct {
	Id      int64 `json:"id" db:"id"`
	OwnerId int64 `json:"owner_id" db:"owner_id"`

	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type DeleteRequest struct {
	Id int64 `json:"id" db:"id"`
}
