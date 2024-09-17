package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	statement, _ := database.Prepare("INSERT INTO users (name, cpf) VALUES (?, ?)")

	response, _ := statement.Exec("John Doe", "07929725545")
	id, _ := response.LastInsertId()
	
	println(id)

	lines, _ := response.RowsAffected()
	println(lines)
}
