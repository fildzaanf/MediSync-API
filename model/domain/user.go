package domain

import (
	"gorm.io/gorm"
	//"github.com/asaskevich/govalidator"
)

type User struct {
	*gorm.Model
	
	Fullname  string     `gorm:"not null" valid:"required~fullname is required"`
	Email     string     `gorm:"not null" valid:"required~email is required, email~invalid email"`
	Password  string     `gorm:"not null" valid:"required~password is required"`
	MedicalID MedicalID  `gorm:"ForeignKey:UserID;references:ID"` // one to one
	Schedules []Schedule `gorm:"ForeignKey:MedicineID;references:ID"` // one to many
}

