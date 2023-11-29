package chatbotDemo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBContent struct {
	ID        int
	Answer    string
	Feedback  string
	CreatedAt string
}

func connectDB() *sql.DB {
	dataSourceName := "root:password@tcp(localhost:3306)/demoSQL"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("DB Connection error %v\n", err)
	}
	return db
}

func SaveContent(content Content) {
	db := connectDB()

	_, err := db.Exec("INSERT INTO demotable (answer, feedback) VALUES (?, ?)", content.Answer, content.Feedback)
	if err != nil {
		fmt.Printf("DB Exec error %v\n", err)
	}
	fmt.Println("Content inserted successfully.")
}

func ShowJSONFromDB() {
	records := ReadDB()
	fmt.Println(StructToJSON(records))
}

func ReadDB() []DBContent {
	db := connectDB()
	// Ping make sure connection to DB
	err := db.Ping()
	if err != nil {
		fmt.Printf("DB Ping error %v\n", err)
	}

	rows, err := db.Query("SELECT * FROM demotable")
	if err != nil {
		fmt.Printf("DB Query error %v\n", err)
	}
	defer rows.Close()

	var records []DBContent

	for rows.Next() {
		var r DBContent
		if err = rows.Scan(&r.ID, &r.Answer, &r.Feedback, &r.CreatedAt); err != nil {
			fmt.Printf("DB Scan error %v\n", err)
		}
		records = append(records, r)
	}
	return records
}

func StructToJSON(records []DBContent) string {
	b, err := json.Marshal(records)
	if err != nil {
		fmt.Printf("JSON Decoding error %v\n", err)
	}
	return string(b)
}
