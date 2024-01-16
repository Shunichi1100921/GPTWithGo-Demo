package main

import (
	"demo/web/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/chat/stream", handlers.HandleStreamChat)
	r.POST("/chat/json", handlers.HandleJSONChat)

	r.GET("/chat/stream/history", handlers.HandleGetStreamChatHistory)
	r.GET("/chat/json/history", handlers.HandleGetJSONChatHistory)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
