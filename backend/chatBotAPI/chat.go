package chatBotAPI

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func createOpenAIClient(apiKey string) *openai.Client {
	return openai.NewClient(apiKey)
}

func getOrCreateChatMessages(chatInput ChatInput, stream bool) (messages []openai.ChatCompletionMessage, err error) {
	chatHistory, err := GetChatHistory(chatInput.ChatID, stream)
	if err != nil {
		return messages, fmt.Errorf("Error getting chat history: %v\n", err)
	}

	messages = createInitialMessage(stream)

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

	return messages, err
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

func createChatRequest(chatInput ChatInput, stream bool) (req openai.ChatCompletionRequest, err error) {
	messages, err := getOrCreateChatMessages(chatInput, stream)
	if err != nil {
		return req, fmt.Errorf("Error getting or creating chat messages: %v\n", err)
	}
	if stream {
		req = openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 1000,
			Messages:  messages,
			Stream:    stream,
		}
		return
	} else {
		req = openai.ChatCompletionRequest{
			Model:          openai.GPT3Dot5Turbo1106,
			MaxTokens:      1000,
			Messages:       messages,
			ResponseFormat: &openai.ChatCompletionResponseFormat{Type: openai.ChatCompletionResponseFormatTypeJSONObject},
		}
		return
	}
}

func CreateChatCompletionStream(chatInput ChatInput) (stream *openai.ChatCompletionStream, err error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := createOpenAIClient(apiKey)
	ctx := context.Background()

	request, err := createChatRequest(chatInput, true)
	stream, err = client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("Error creating Stream chat completion: %v\n", err)
	}
	return
}

func CreateChatCompletionJSON(chatInput ChatInput) (response openai.ChatCompletionResponse, err error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := createOpenAIClient(apiKey)
	ctx := context.Background()

	request, err := createChatRequest(chatInput, false)
	if err != nil {
		return response, fmt.Errorf("Error creating JSON chat completion request: %v\n", err)
	}
	response, err = client.CreateChatCompletion(ctx, request)
	if err != nil {
		return response, fmt.Errorf("Error creating JSON chat completion: %v\n", err)
	}
	return
}
