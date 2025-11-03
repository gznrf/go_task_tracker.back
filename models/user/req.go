package m_user

type GetRequest struct {
}
type GetByIdRequest struct {
	Id int64 `json:"id" db:"id"`
}
type UpdateRequest struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}
type DeleteRequest struct {
	Id int64 `json:"id" db:"id"`
}
