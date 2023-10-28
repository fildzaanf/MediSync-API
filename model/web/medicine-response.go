package web

type MedicineResponse struct {
	ID          int               `json:"id" form:"id"`
	Name        string             `json:"name" form:"name"`
	Amount      int                `json:"amount" form:"amount"`
	Details     string             `json:"details" form:"details"`
	BatchNumber int                `json:"batch_number" form:"batch_number"`
	Category    []CategoryResponse `json:"category" form:"category"` // one to many
}
