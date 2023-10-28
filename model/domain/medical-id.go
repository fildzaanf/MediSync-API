package domain

import (
	"gorm.io/gorm"
)

type MedicalID struct {
	*gorm.Model

	UserID           uint   `json:"user_id" form:"user_id"`
	Birthdate        string `json:"birthdate" form:"birthdate"`
	Gender           string `json:"gender" form:"gender" gorm:"type:enum('male', 'female')"`
	BloodType        string `json:"blood_type" form:"blood_type" gorm:"type:enum('A', 'B', 'O', 'AB')"`
	Height           int    `json:"height" form:"height"`
	Weight           int    `json:"weight" form:"weight"`
	MedicalCondition string `json:"medical_condition" form:"medical_condition"`
	EmergencyContact string `json:"emergency_contact" form:"emergency_contact"`
}
