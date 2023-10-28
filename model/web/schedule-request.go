package web

type ScheduleRequest struct {
	Name    string `json:"name" form:"name"`
	Details string `json:"details" form:"details"`
	Time    string `json:"time" form:"time"`
	Repeat  string `json:"repeat" form:"repeat"`
	Status  string `json:"status" form:"status"`
}