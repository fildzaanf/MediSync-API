package controller

import (
	"app/model/web"
	"testing"
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/sashabaranov/go-openai"
)

func TestOpenAIMediChatResponseController(t *testing.T) {

	request := &web.MediChatRequest{
		Message: "Apa obat yang aman untuk flu?",
	}
	response, err := OpenAIMediChatResponseController(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Message)
	assert.NotEmpty(t, response.Response)
	assert.Contains(t, response.Response, "Obat untuk flu yang umumnya aman adalah panadol")
	assert.GreaterOrEqual(t, len(response.Response), 10)
}

func TestGetCompletionFromMessage(t *testing.T) {
	ctx := context.Background()
	client := openai.NewClient("OPENAI_KEY") 
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "System message",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "User message",
		},
	}
	model := openai.GPT3Dot5Turbo
	response, err := GetCompletionFromMessage(ctx, client, messages, model)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Choices[0].Message.Content)
	expectedResponse := "Expected response from the model"
	assert.Equal(t, expectedResponse, response.Choices[0].Message.Content)
}