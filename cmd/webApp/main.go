package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"demo/web/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/chat/stream", handlers.HandleStreamChat)
	r.POST("/chat/json", handlers.HandleJSONChat)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
