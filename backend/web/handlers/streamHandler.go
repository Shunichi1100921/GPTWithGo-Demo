package handlers

import (
	"demo/chatBotAPI"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"strconv"
)

func HandleStreamChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	var chatInput chatBotAPI.ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bind JSON error": err.Error()})
		return
	}

	stream, err := chatBotAPI.CreateChatCompletionStream(chatInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Creating Stream Chat Completion error": err.Error()})
		return
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
	if err := chatBotAPI.SaveChatHistory(chatInput, finalResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Saving chat history error": err.Error()})
		return
	}
}

func HandleGetStreamChatHistory(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid ID error": err.Error()})
		return
	}
	history, err := chatBotAPI.GetChatHistory(id, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Getting Chat History error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}
