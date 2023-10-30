package domain

import (
	"gorm.io/gorm"
)

type Medicine struct {
	*gorm.Model

	ScheduleID  uint
	CategoryID  uint 	
	Name        string
	Amount      int
	Details     string
	BatchNumber int
	// Category    Category `gorm:"ForeignKey:CategoryID;references:ID"`
}
