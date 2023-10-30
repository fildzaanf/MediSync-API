package web

type CategoryRequest struct {
	Name    string `json:"name" form:"name"`
	Details string `json:"details" form:"details"`
}
