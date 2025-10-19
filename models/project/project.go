package m_project

type Project struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	OwnerId     int64  `json:"owner_id" db:"owner_id"`
}
