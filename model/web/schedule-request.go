package web

type ScheduleRequest struct {
	Name    string `json:"name" form:"name"`
	Details string `json:"details" form:"details"`
	Minute  string   `json:"minute" form:"minute"` // 0-59
	Hour    string    `json:"hour" form:"hour"`     // 0-23
	Day     string    `json:"day" form:"day"`       // 0 or 7 (sunday); 1 (monday)
	// Email   string `json:"email" form:"email"`
	Subject string `json:"subject" form:"subject"`
	Body    string `json:"body" form:"body"`
	// Status  string `json:"status" form:"status"`
}
