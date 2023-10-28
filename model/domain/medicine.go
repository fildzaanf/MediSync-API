package domain

import (
	"gorm.io/gorm"
)

type Medicine struct {
	*gorm.Model

	ScheduleID  uint   `json:"schedule_id" form:"schedule_id"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	Name        string `json:"name" form:"name"`
	Amount      int    `json:"amount" form:"amount"`
	Details     string `json:"details" form:"details"`
	BatchNumber int    `json:"batch_number" form:"batch_number"`
}
