package domain

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Fullname  string     `json:"fullname" form:"fullname" validate:"required"`
	Email     string     `json:"email" form:"email" validate:"required"`
	Password  string     `json:"password" form:"password" validate:"required"`
	MedicalID MedicalID  `gorm:"foreignKey:UserID;references:ID;association_autoupdate:true;association_autocreate:false"` // one to one
	// Schedules []Schedule `gorm:"foreignKey:UserID;references:ID;association_autoupdate:true;association_autocreate:false"` // one to many

}
