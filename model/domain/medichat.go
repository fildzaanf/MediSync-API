package domain

import (
	"gorm.io/gorm"
)

type MediChat struct {
	*gorm.Model

	Message  string `json:"message" form:"message"`
	Response string `json:"response" form:"response"`
}
