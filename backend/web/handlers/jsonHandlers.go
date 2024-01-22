package handlers

import (
	"demo/chatBotAPI"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleJSONChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	var chatInput chatBotAPI.ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"BindJSON error": err.Error()})
		return
	}

	response, err := chatBotAPI.CreateChatCompletionJSON(chatInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Creating JSON Chat Completion error": err.Error()})
		return
	}

	msg := response.Choices[0].Message.Content
	msgByte := []byte(msg)

	var content chatBotAPI.JSONChatResponse
	err = json.Unmarshal(msgByte, &content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"JSON Unmarshal error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, content)

	// Save the chat history to the database.
	if err = chatBotAPI.SaveChatHistoryWithFeedback(chatInput, content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Saving Chat history error": err.Error()})
		return
	}
}

func HandleGetJSONChatHistory(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid ID error.": err.Error()})
		return
	}
	history, err := chatBotAPI.GetChatHistory(id, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Getting chat history error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}
