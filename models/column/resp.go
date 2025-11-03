package m_column

type CreateResponse struct {
	CreatedId int64 `json:"created_id"`
}
type GetResponse struct {
	ColumnsList []Column `json:"columns_list"`
}
type GetByBoardIdResponse struct {
	ColumnsList []Column `json:"columns_list"`
}
type GetByIdResponse struct {
	Column Column `json:"column"`
}
type UpdateResponse struct {
	UpdatedId int64 `json:"updated_id"`
}
type DeleteResponse struct {
	DeletedId int64 `json:"deleted_id"`
}
