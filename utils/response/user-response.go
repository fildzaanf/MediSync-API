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
		}
		results = append(results, userCreateResponse)
	}

	return results
}

func ConvertToGetUser(user *domain.User) web.UserCreateResponse {
	return web.UserCreateResponse{
		ID:       int(user.ID),
		Fullname:     user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}
}
