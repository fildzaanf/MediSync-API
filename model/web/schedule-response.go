package web

import "app/model/domain"

type ScheduleResponse struct {
	ID         int             `json:"id" form:"id"`
	MedicineID int             `json:"medicine_id" form:"medicine_id"`
	Name       string          `json:"name" form:"name"`
	Details    string          `json:"details" form:"details"`
	Minute     string          `json:"minute" form:"minute"` // 0-59
	Hour       string          `json:"hour" form:"hour"`     // 0-23
	Day        string          `json:"day" form:"day"`       // 0 or 7 (sunday); 1 (monday)
	Subject    string          `json:"subject" form:"subject"`
	// Body       string          `json:"body" form:"body"`
	Medicine   domain.Medicine `json:"medicine" form:"medicine"` // one to many
}
