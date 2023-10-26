package web

type UserResponse struct {
	ID        int                `json:"id" form:"id"`
	Fullname  string             `json:"fullname" form:"fullname"`
	Email     string             `json:"email" form:"email"`
	Password  string             `json:"password" form:"password"`
	MedicalID MedicalIDResponse  `json:"medical_id" form:"medical_id"` // one to one
	Schedules []ScheduleResponse `json:"schedules" form:"schedules"`   // one to many

}

type UserCreateResponse struct {
	ID        int                `json:"id" form:"id"`
	Fullname  string             `json:"fullname" form:"fullname"`
	Email     string             `json:"email" form:"email"`
	Password  string             `json:"password" form:"password"`
	MedicalID MedicalIDResponse  `json:"medical_id" form:"medical_id"` // one to one
	Schedules []ScheduleResponse `json:"schedules" form:"schedules"`   // one to many

}

type UserLoginResponse struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}
