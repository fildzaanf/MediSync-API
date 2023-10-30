package domain

import (
	"gorm.io/gorm"
)

type Schedule struct {
	*gorm.Model

	// MedicineID uint
	UserID   uint
	Name     string
	Details  string
	Minute   string // 0-59
	Hour     string // 0-23
	Day      string // 0 or 7 (sunday); 1 (monday)
	Email    string
	Subject  string
	Body     string
	Medicine Medicine `gorm:"ForeignKey:ScheduleID;references:ID"` // one to many

}
