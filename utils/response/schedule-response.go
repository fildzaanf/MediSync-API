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
			Minute:  schedule.Minute,
			Hour:    schedule.Hour,
			// Date:    schedule.Date,
			// Month:   schedule.Month,
			Day:     schedule.Day,
			// Email:   schedule.Email,
			Subject: schedule.Subject,
			// Body:    schedule.Body,
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
		Minute:  schedule.Minute,
		Hour:    schedule.Hour,
		//Date:    schedule.Date,
		//Month:   schedule.Month,
		Day:     schedule.Day,
		// Email:   schedule.Email,
		Subject: schedule.Subject,
		Body:    schedule.Body,
	}
}
