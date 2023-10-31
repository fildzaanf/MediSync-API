package domain

import (
	"gorm.io/gorm"
)

type MedicalID struct {
	*gorm.Model

	UserID           uint 
	Birthdate        string
	Gender           string `gorm:"type:enum('male', 'female')"`
	BloodType        string `gorm:"type:enum('A', 'B', 'O', 'AB')"`
	Height           int
	Weight           int
	MedicalCondition string
	EmergencyContact string
}
