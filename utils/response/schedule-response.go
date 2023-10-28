package response

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToGetAllSchedules(schedules []domain.Schedule) []web.ScheduleResponse {
	var results []web.ScheduleResponse
	for _, schedule := range schedules {
		scheduleResponse := web.ScheduleResponse{
			ID:      int(schedule.ID),
			Name:    schedule.Name,
			Details: schedule.Details,
			Time:    schedule.Time,
			Repeat:  schedule.Repeat,
			Status:  schedule.Status,
		}
		results = append(results, scheduleResponse)
	}
	return results
}

func ConvertToGetSchedule(schedule *domain.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID:      int(schedule.ID),
		Name:    schedule.Name,
		Details: schedule.Details,
		Time:    schedule.Time,
		Repeat:  schedule.Repeat,
		Status:  schedule.Status,
	}
}
