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
			MedicineID: int(schedule.MedicineID),
			Name:    schedule.Name,
			Details: schedule.Details,
			Minute:  schedule.Minute,
			Hour:    schedule.Hour,
			Day:     schedule.Day,
			// Email:   schedule.Email,
			Subject: schedule.Subject,
			Medicine: schedule.Medicine,
			// Body:    schedule.Body,
		}
		results = append(results, scheduleResponse)
	}
	return results
}

func ConvertToCreateSchedule(schedule *domain.Schedule, medicine domain.Medicine) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID:      int(schedule.ID),
		MedicineID: int(schedule.MedicineID),
		Name:    schedule.Name,
		Details: schedule.Details,
		Minute:  schedule.Minute,
		Hour:    schedule.Hour,
		Day:     schedule.Day,
		Subject: schedule.Subject,
		Medicine: medicine,
	}
}

func ConvertToGetSchedule(schedule *domain.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID:      int(schedule.ID),
		MedicineID: int(schedule.MedicineID),
		Name:    schedule.Name,
		Details: schedule.Details,
		Minute:  schedule.Minute,
		Hour:    schedule.Hour,
		Day:     schedule.Day,
		Subject: schedule.Subject,
		Medicine: schedule.Medicine,
	}
}
