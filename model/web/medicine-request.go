package web

type MedicineRequest struct {
	CategoryID  int    `json:"category_id" form:"category_id"`
	Name        string `json:"name" form:"name"`
	Amount      int    `json:"amount" form:"amount"`
	Details     string `json:"details" form:"details"`
	BatchNumber int    `json:"batch_number" form:"batch_number"`
}
