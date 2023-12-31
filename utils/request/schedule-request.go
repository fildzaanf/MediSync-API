package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToScheduleRequest(schedule web.ScheduleRequest) *domain.Schedule {
	return &domain.Schedule{
		Name:    schedule.Name,
		Details: schedule.Details,
		Minute:  schedule.Minute,
		Hour:    schedule.Hour,
		Day:     schedule.Day,
		Email:   schedule.Email,
		Subject: schedule.Subject,
	}
}
