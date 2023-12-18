package chatBotAPI

import (
	"database/sql"
	"log"
)

func connectDB() *sql.DB {
	dataSourceName := "root:password@tcp(localhost:3306)/demoSQL?parseTime=True"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("DB Connection error %v\n", err)
	}
	return db
}

func getChatHistory(chatInput ChatInput) []ChatHistory {
	db := connectDB()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error %v\n", err)
	}

	rows, err := db.Query("SELECT user_prompt, bot_response FROM demoSQL.chatHistory WHERE id = ? ORDER BY created_at", chatInput.ChatID)
	if err != nil {
		log.Fatalf("DB Query error %v\n", err)
	}
	defer rows.Close()

	var chatHistories []ChatHistory

	for rows.Next() {
		var h ChatHistory
		if err = rows.Scan(&h.UserPrompt, &h.BotResponse); err != nil {
			log.Fatalf("DB Scan error %v\n", err)
		}
		chatHistories = append(chatHistories, h)
	}
	return chatHistories
}

func SaveChatHistory(chatInput ChatInput, finalResponse string) {
	db := connectDB()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error %v\n", err)
	}

	_, err := db.Exec("INSERT INTO demoSQL.chatHistory (id, user_prompt, bot_response) VALUES (?, ?, ?)", chatInput.ChatID, chatInput.Message, finalResponse)
	if err != nil {
		log.Fatalf("DB Exec error %v\n", err)
	}
	log.Println("Chat history was saved successfully.")
}

func getChatWithFeedbackHistory(chatInput ChatInput) []ChatHistory {
	db := connectDB()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error %v\n", err)
	}

	rows, err := db.Query("SELECT user_prompt, bot_response FROM demoSQL.chatWithFeedbackHistory WHERE id = ? ORDER BY created_at", chatInput.ChatID)
	if err != nil {
		log.Fatalf("DB Query error %v\n", err)
	}
	defer rows.Close()

	var chatHistories []ChatHistory

	for rows.Next() {
		var h ChatHistory
		if err = rows.Scan(&h.UserPrompt, &h.BotResponse); err != nil {
			log.Fatalf("DB Scan error %v\n", err)
		}
		chatHistories = append(chatHistories, h)
	}
	return chatHistories
}

func SaveChatHistoryWithFeedback(chatInput ChatInput, response JSONChatResponse) {
	db := connectDB()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error %v\n", err)
	}

	_, err := db.Exec("INSERT INTO demoSQL.chatWithFeedbackHistory (id, user_prompt, bot_response, feedback) VALUES (?, ?, ?, ?)", chatInput.ChatID, chatInput.Message, response.Answer, response.Feedback)
	if err != nil {
		log.Fatalf("DB Exec error %v\n", err)
	}
	log.Println("Chat history was saved successfully.")
}
