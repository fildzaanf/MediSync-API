package response

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToMediChatResponse(medichat *domain.MediChat) *web.MediChatResponse {
	return &web.MediChatResponse{
		Response: medichat.Response,
	}
}

