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

	transaction, _ := database.Begin()
	statement, _ := transaction.Prepare("INSERT INTO users (id, name, cpf) VALUES (?, ?, ?)")

	statement.Exec(102, "John Doe", "07929725545")
	statement.Exec(103, "Mary Doe", "12345678901")

	_, err = statement.Exec(102, "James Doe", "07929725545")

	if err != nil {
		transaction.Rollback()
		log.Fatal(err)
	}

	transaction.Commit()
}
