package domain

import (
	"gorm.io/gorm"
)

type MediChat struct {
	*gorm.Model

	Message  string 
	Response string 
}
