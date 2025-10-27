package m_column

type CreateRequest struct {
	Name    string `json:"name" db:"name"`
	BoardId int64  `json:"board_id" db:"board_id"`
}
type GetRequest struct {
}
type GetByBoardIdRequest struct {
	BoardId int64 `json:"board_id" db:"board_id"`
}
type GetByIdRequest struct {
	ColumnId int64 `json:"column_id" db:"column_id"`
}
type UpdateRequest struct {
	ColumnId int64  `json:"column_id" db:"column_id"`
	Name     string `json:"name" db:"name"`
	BoardId  int64  `json:"board_id" db:"board_id"`
}
type DeleteRequest struct {
	ColumnId int64 `json:"column_id" db:"column_id"`
}
