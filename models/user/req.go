package m_user

type GetUserByIdRequest struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	UserId   int64  `json:"id"`
	Username string `json:"name" db:"name"`
}

type DeleteUserRequest struct {
	Token string `json:"token"`
}
