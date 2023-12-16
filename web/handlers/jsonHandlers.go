package handlers

import (
	"demo/chatBotAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleJSONChat(c *gin.Context) {
	var chatInput chatBotAPI.ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
