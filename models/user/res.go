package m_user

type GetResponse struct {
	UsersList []User `json:"users_list"`
}
type GetByIdResponse struct {
	User User `json:"user"`
}
type UpdateResponse struct {
	UpdatedId int64 `json:"updated_id"`
}
type DeleteResponse struct {
	DeletedId int64 `json:"deleted_id"`
}
