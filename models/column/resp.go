package m_column

type CreateResponse struct {
}
type GetResponse struct {
}
type GetByBoardIdResponse struct {
}
type GetByIdResponse struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	BoardId int    `json:"board_id" db:"board_id"`
}
type UpdateResponse struct {
}
type DeleteResponse struct {
}
