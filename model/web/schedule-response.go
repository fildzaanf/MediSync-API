package web

type ScheduleResponse struct {
	ID      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Details string `json:"details" form:"details"`
	Minute  string `json:"minute" form:"minute"` // 0-59
	Hour    string `json:"hour" form:"hour"`     // 0-23
	// Date    string `json:"date" form:"date"`     // 1-31
	// Month   string `json:"month" form:"month"`   // 1-12
	Day     string `json:"day" form:"day"`       // 0 or 7 (sunday); 1 (monday)
	// Email   string `json:"email" form:"email"`
	Subject string `json:"subject" form:"subject"`
	Body    string `json:"body" form:"body"`
	// Status   string             `json:"status" form:"status"`
	// Message  string             `json:"message" form:"message"`
	Medicine []MedicineResponse `json:"medicine" form:"medicine"` // one to many
}
