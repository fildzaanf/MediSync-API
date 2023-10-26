package web

type CategoryResponse struct {
	ID       uint               `json:"id" form:"id"`
	Name     string             `json:"name" form:"name"`
	Details  string             `json:"details" form:"details"`
	Medicine []MedicineResponse `json:"medicine" form:"medicine"` // one to many
}
