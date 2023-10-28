package web

type MediChatRequest struct {
	Message string `json:"message" form:"message"`
}
