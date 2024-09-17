package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	database, err := sql.Open("mysql", "root:root@/")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	exec(database, "CREATE DATABASE IF NOT EXISTS gocourse")
	exec(database, "USE gocourse")
	exec(database, "DROP TABLE IF EXISTS users")
	exec(database, `CREATE TABLE users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(80),
		cpf CHAR(11)
	)`)
}
