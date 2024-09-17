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

	statement, _ := database.Prepare("UPDATE users SET name = ?, cpf = ? WHERE id = ?")
	statement.Exec("John Doe edited", "07929725545", 1)
}
