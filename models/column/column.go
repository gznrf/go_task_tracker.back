package m_column

type Column struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	BoardId int    `json:"board_id" db:"board_id"`
}
