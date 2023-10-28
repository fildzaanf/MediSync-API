package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToScheduleRequest(schedule web.ScheduleRequest) *domain.Schedule {
	return &domain.Schedule{
		Name:    schedule.Name,
		Details: schedule.Details,
		Time:    schedule.Time,
		Repeat:  schedule.Repeat,
		Status:  schedule.Status,
	}
}
