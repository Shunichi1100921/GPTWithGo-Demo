package chatbotDemo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DBContent struct {
	ID        int
	Answer    string
	Feedback  string
	CreatedAt time.Time
}

func ReadContent(content Content) {
	dataSourceName := "root:password@tcp(localhost:3306)/demoSQL"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("DB Connection error %v\n", err)
	}
	defer db.Close()

	// Ping make sure connection to DB
	err = db.Ping()
	if err != nil {
		fmt.Printf("DB Ping error %v\n", err)
	}

	rows, err := db.Query("SELECT * FROM demotable")
	if err != nil {
		fmt.Printf("DB Query error %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			answer   string
			feedback string
		)
		if err = rows.Scan(&answer, &feedback); err != nil {
			fmt.Printf("DB Scan error %v\n", err)
		}
		fmt.Printf("")
	}

}

func ShowJSON() {
	c := Content{"answer", "feadback"}

	b, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("JSON Decoding error %v\n", err)
	}

	fmt.Println(b)
}

func SaveContent(content Content) {
	dataSourceName := "root:password@tcp(localhost:3306)/demoSQL"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("DB Connection error %v\n", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO demotable (answer, feedback) VALUES (?, ?)", content.Answer, content.Feedback)
	if err != nil {
		fmt.Printf("DB Exec error %v\n", err)
	}
	fmt.Println("Content inserted successfully.")
}
