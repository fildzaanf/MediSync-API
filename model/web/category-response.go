package web

import (
	"app/model/domain"
)

// import "app/model/domain"

type CategoryResponse struct {
	ID       int               `json:"id" form:"id"`
	Name     string            `json:"name" form:"name"`
	Details  string            `json:"details" form:"details"`
	Medicine []domain.Medicine `json:"medicine" form:"medicine"` // one to many
}

