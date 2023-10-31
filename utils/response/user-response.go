package response

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToGetAllUsers(users []domain.User) []web.UserCreateResponse {
	var results []web.UserCreateResponse
	for _, user := range users {
		userCreateResponse := web.UserCreateResponse{
			ID:       int(user.ID),
			Fullname: user.Fullname,
			Email:    user.Email,
			Password: user.Password,
			MedicalID: user.MedicalID,
		}
		for _, schedule := range user.Schedules{
			scheduleResponse := domain.Schedule{
				MedicineID: schedule.MedicineID,
				Name: schedule.Name,
				Details: schedule.Details,
				Minute: schedule.Minute,
				Hour: schedule.Hour,
				Day: schedule.Day,
				Email: schedule.Email,
				Subject: schedule.Subject,
			}
			userCreateResponse.Schedules = append(userCreateResponse.Schedules, scheduleResponse)
		}
 
		results = append(results, userCreateResponse)
	}

	return results
}

func ConvertToGetUser(user *domain.User) web.UserCreateResponse {
	return web.UserCreateResponse{
		ID:       int(user.ID),
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
		MedicalID: user.MedicalID,
		Schedules: user.Schedules,
	}
}
