package handlers

import (
	"demo/chatBotAPI"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

func HandleStreamChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	var chatInput chatBotAPI.ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stream, err := chatBotAPI.CreateChatCompletionStream(chatInput)
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
