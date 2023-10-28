package web

type CategoryRequest struct {
	Name    string `json:"name" form:"name" validate:"required"`
	Details string `json:"details" form:"details"`
}
