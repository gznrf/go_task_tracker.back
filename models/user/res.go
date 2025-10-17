package m_user

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type GetUserByIdResponse struct {
	User User `json:"user"`
}

type UpdateUserResponse struct {
	IsUpdated bool `json:"is_updated"`
}

type DeleteUserResponse struct {
	IsDeleted bool `json:"is_deleted"`
}
