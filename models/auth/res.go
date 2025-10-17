package m_auth

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token         string `json:"token"`
	CreatedUserId int64  `json:"created_user_id"`
}

type ChangePasswordResponse struct {
	Id           int64  `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}
