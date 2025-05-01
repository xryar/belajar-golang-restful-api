package web

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}
