package main

import (
	"demo/chatbotDemo"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/chat/json", func(c *gin.Context) {
		chatbotDemo.ChatJSON()
		response := chatbotDemo.ShowJSONFromDB
		c.JSON(200, response)
	})

	r.GET("chat/stream", func(c *gin.Context) {
		chatbotDemo.ChatStream()
	})

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
