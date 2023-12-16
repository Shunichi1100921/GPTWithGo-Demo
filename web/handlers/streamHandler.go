package handlers

import (
	"context"
	"demo/chatBotAPI"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
)

func HandleStreamChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := chatBotAPI.CreateOpenAIClient(apiKey)
	ctx := context.Background()

	var chatInput chatBotAPI.ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messages := chatBotAPI.GetOrCreateChatMessages(chatInput)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: chatInput.Message,
	})

	request := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Messages:  messages,
		Stream:    true,
	}

	stream, err := client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var finalResponse string

	c.Stream(func(w io.Writer) bool {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return false
		}
		if err != nil {
			return false
		}

		responseMessage := response.Choices[0].Delta.Content

		c.SSEvent("message", responseMessage)
		finalResponse += responseMessage

		return true
	})
	// Save the chat history to the database.
	chatBotAPI.SaveChatHistory(chatInput, finalResponse)
}
