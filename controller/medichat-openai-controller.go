package controller

import (
	"context"
	"os"
	"log"
	"app/model/web"
	"app/model/domain"
	"github.com/joho/godotenv"

	openai "github.com/sashabaranov/go-openai"
)

// Request response to OPENAI
func OpenAIMediChatResponseController(request *web.MediChatRequest) (*domain.MediChat, error) {

	_, err := os.Stat(".env")
    if err == nil {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Failed to fetch .env file")
        }
    }

	ctx := context.Background()
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))
	model := openai.GPT3Dot5Turbo

	message := []openai.ChatCompletionMessage{{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Anda adalah AI MediChat. Anda dirancang untuk memberikan informasi dan jawaban terkait topik kesehatan, terutama pertanyaan tentang obat-obatan.",
	},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: request.Message,
		},
	}

	AIResponse, err := GetCompletionFromMessage(ctx, client, message, model)
	if err != nil {
		return &domain.MediChat{
			Message: request.Message, 
			Response: "Fail To Get Completion",
		}, err
	}

	answer := AIResponse.Choices[0].Message.Content
	response := domain.MediChat{
		Message: request.Message,
		Response: answer,
	}

	return &response, nil
}

// Generate Response 
func GetCompletionFromMessage(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	response, err := client.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	return response, err
}
