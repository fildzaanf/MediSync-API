package web

type MedicineResponse struct {
	ID          uint               `json:"id" form:"id"`
	Name        string             `json:"name" form:"name"`
	Amount      int                `json:"amount" form:"amount"`
	Details     string             `json:"details" form:"details"`
	BatchNumber int                `json:"batch_number" form:"batch_number"`
	ExpiredAt   string             `json:"expired_at" form:"expired_at"`
	Category    []CategoryResponse `json:"category" form:"category"` // one to many
}
