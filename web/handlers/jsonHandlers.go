package handlers

import (
	"bytes"
	"demo/chatBotAPI"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleJSONChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	var chatInput chatBotAPI.ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := chatBotAPI.CreateChatCompletionJSON(chatInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	msg := response.Choices[0].Message.Content
	msgByte := []byte(msg)

	var content chatBotAPI.JSONChatResponse
	err = json.Unmarshal(msgByte, &content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, msgByte, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	prettyContent := buf.String()

	c.JSON(http.StatusOK, prettyContent)

	// Save the chat history to the database.
	chatBotAPI.SaveChatHistoryWithFeedback(chatInput, content)
}
