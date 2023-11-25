package chatbotDemo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

func main() {
	demoContent := Content{Answer: "Answer", Feedback: "Feedback"}
	SaveContent(demoContent)
}
