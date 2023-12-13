package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleStreamChat(c *gin.Context) {
	var chatInput ChatInput
	if err := c.BindJSON(&chatInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("StreamChat")
	fmt.Println(chatInput.Message)
}
