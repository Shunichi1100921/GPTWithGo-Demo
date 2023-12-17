package chatBotAPI

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func createOpenAIClient(apiKey string) *openai.Client {
	return openai.NewClient(apiKey)
}

func getOrCreateChatMessages(chatInput ChatInput) []openai.ChatCompletionMessage {
	chatHistory := getChatHistory(chatInput)

	messages := createInitialMessage()

	for _, h := range chatHistory {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: h.UserPrompt,
		})
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: h.BotResponse,
		})
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: chatInput.Message,
	})

	return messages
}

func createInitialMessage() []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Answer my question.",
		},
	}
}

func createChatRequest(chatInput ChatInput) openai.ChatCompletionRequest {
	messages := getOrCreateChatMessages(chatInput)
	return openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Messages:  messages,
		Stream:    true,
	}
}

func CreateChatCompletionStream(chatInput ChatInput) (stream *openai.ChatCompletionStream, err error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := createOpenAIClient(apiKey)
	ctx := context.Background()

	request := createChatRequest(chatInput)
	stream, err = client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		log.Fatalf("Error creating chat completion: %v\n", err)
	}
	return
}
