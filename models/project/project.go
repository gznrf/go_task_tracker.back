package m_project

type Project struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	OwnerId     int    `json:"owner_id" db:"owner_id"`
}
