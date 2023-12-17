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

func getOrCreateChatMessages(chatInput ChatInput, stream bool) []openai.ChatCompletionMessage {
	chatHistory := getChatHistory(chatInput)

	messages := createInitialMessage(stream)

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

func createInitialMessage(stream bool) []openai.ChatCompletionMessage {
	var firstPrompt string
	if stream {
		firstPrompt = "Answer my questions."
	} else {
		firstPrompt = "You're the boss of me who answers my questions.  You gotta give me the answers and feedback " +
			"about the way I asked. You should return in JSON format, which has two keys.  One is the answer and " +
			"the other is feedback."
	}
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: firstPrompt,
		},
	}
}

func createChatRequest(chatInput ChatInput, stream bool) openai.ChatCompletionRequest {
	messages := getOrCreateChatMessages(chatInput, stream)
	if stream {
		return openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 1000,
			Messages:  messages,
			Stream:    stream,
		}
	} else {
		return openai.ChatCompletionRequest{
			Model:          openai.GPT3Dot5Turbo1106,
			MaxTokens:      1000,
			Messages:       messages,
			ResponseFormat: &openai.ChatCompletionResponseFormat{Type: openai.ChatCompletionResponseFormatTypeJSONObject},
		}
	}
}

func CreateChatCompletionStream(chatInput ChatInput) (stream *openai.ChatCompletionStream, err error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := createOpenAIClient(apiKey)
	ctx := context.Background()

	request := createChatRequest(chatInput, true)
	stream, err = client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		log.Fatalf("Error creating chat completion: %v\n", err)
	}
	return
}

func CreateChatCompletionJSON(chatInput ChatInput) (response openai.ChatCompletionResponse, err error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := createOpenAIClient(apiKey)
	ctx := context.Background()

	request := createChatRequest(chatInput, false)
	response, err = client.CreateChatCompletion(ctx, request)
	if err != nil {
		log.Fatalf("Error creating chat completion: %v\n", err)
	}
	return
}
