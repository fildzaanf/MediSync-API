package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model

	Name     string
	Details  string
	Medicine []Medicine `gorm:"ForeignKey:CategoryID;references:ID"` // one to many
}
