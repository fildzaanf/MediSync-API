package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model

	Name     string     `json:"name" form:"name"`
	Details  string     `json:"details" form:"details"`
	Medicine []Medicine // one to many
}
