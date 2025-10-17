package m_auth

type LoginRequest struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}

type RegisterRequest struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}
