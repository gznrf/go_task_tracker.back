package m_auth

type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}
type RegisterRequest struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}
