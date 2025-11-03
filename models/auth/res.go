package m_auth

type LoginResponse struct {
	UserId int64 `json:"user_id"`
}
type RegisterResponse struct {
	CreatedId int64 `json:"created_id"`
}
