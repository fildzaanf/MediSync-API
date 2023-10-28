package domain

import (
	"gorm.io/gorm"
)

type Schedule struct {
	*gorm.Model

	MedicineID uint       `json:"medicine_id" form:"medicine_id"`
	// UserID     uint       `json:"user_id" form:"user_id"`
	Name       string     `json:"name" form:"name"`
	Details    string     `json:"details" form:"details"`
	Time       string     `json:"time" form:"time"`
	Repeat     string     `json:"repeat" form:"repeat"`
	Status     string     `json:"status" form:"status"`
	Medicine   []Medicine // one to many

}
