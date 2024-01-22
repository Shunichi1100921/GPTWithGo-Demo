package chatBotAPI

import (
	"database/sql"
	"fmt"
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

func GetChatHistory(chatID int, stream bool) []ChatHistory {
	db := connectDB()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error %v\n", err)
	}

	var sqlStatement string
	if stream {
		sqlStatement = "SELECT user_prompt, bot_response FROM demoSQL.chatHistory WHERE id = ? ORDER BY created_at"
	} else {
		sqlStatement = "SELECT user_prompt, bot_response FROM demoSQL.chatWithFeedbackHistory WHERE id = ? ORDER BY created_at"
	}

	rows, err := db.Query(sqlStatement, chatID)
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

func SaveChatHistory(chatInput ChatInput, finalResponse string) (err error) {
	db := connectDB()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error %v\n", err)
	}

	_, err = db.Exec("INSERT INTO demoSQL.chatHistory (id, user_prompt, bot_response) VALUES (?, ?, ?)", chatInput.ChatID, chatInput.Message, finalResponse)
	if err != nil {
		return fmt.Errorf("DB Exec error in SaveChatHistory %v\n", err)
	}
	log.Println("Chat history was saved successfully.")
	return
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
