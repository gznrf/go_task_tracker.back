package m_board

type Board struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ProjectId int64  `json:"project_id" db:"project_id"`
}
