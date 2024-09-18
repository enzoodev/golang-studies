package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id int
	name string
	cpf string
}

func main() {
	database, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	rows, _ := database.Query("SELECT id, name, cpf FROM users WHERE id > ?", 1)
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name, &u.cpf)
		fmt.Println(u)
	}
}
