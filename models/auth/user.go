package m_auth

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"email" db:"username"`
	Password string `json:"password" db:"password"`
}
