package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToUserCreateRequest(user web.UserCreateRequest) *domain.User {
	return &domain.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}
}
