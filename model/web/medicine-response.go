package web

import (
	// "app/model/domain"
)

type MedicineResponse struct {
	ID          int             `json:"id" form:"id"`
	ScheduleID  int				`json:"schedule_id" form:"schedule_id"`
	CategoryID  int             `json:"category_id" form:"category_id"`
	Name        string          `json:"name" form:"name"`
	Amount      int             `json:"amount" form:"amount"`
	Details     string          `json:"details" form:"details"`
	BatchNumber int             `json:"batch_number" form:"batch_number"`
	// Category    domain.Category `json:"category" form:"category"` // one to many
}
