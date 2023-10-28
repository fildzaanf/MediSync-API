package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToMediChatRequest(medichat web.MediChatRequest) *domain.MediChat {
	return &domain.MediChat{
		Message: medichat.Message,
	}
}
