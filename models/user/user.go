package m_user

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Password string `json:"password" db:"password_hash"`
}
