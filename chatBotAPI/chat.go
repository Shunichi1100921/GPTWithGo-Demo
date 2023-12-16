package chatBotAPI

import (
	"github.com/sashabaranov/go-openai"
)

func CreateOpenAIClient(apiKey string) *openai.Client {
	return openai.NewClient(apiKey)
}

func GetOrCreateChatMessages(chatInput ChatInput) []openai.ChatCompletionMessage {
	chatHistory := getChatHistory(chatInput)
	if chatHistory == nil {
		return createInitialMessage()
	}
	var messages []openai.ChatCompletionMessage

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
