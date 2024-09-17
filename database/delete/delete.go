package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	statement, _ := database.Prepare("DELETE FROM users WHERE id = ?")
	statement.Exec(102)
}
