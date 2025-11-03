package m_project

type CreateResponse struct {
	CreatedId int64 `json:"created_id"`
}
type GetResponse struct {
	ProjectsList []Project `json:"projects_list"`
}
type GetByUserIdResponse struct {
	ProjectsList []Project `json:"projects_list"`
}
type GetByIdResponse struct {
	Project Project `json:"project"`
}
type UpdateResponse struct {
	UpdatedId int64 `json:"updated_id"`
}
type DeleteResponse struct {
	DeletedId int64 `json:"deleted_id"`
}
