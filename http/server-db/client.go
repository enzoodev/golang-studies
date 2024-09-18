package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Cpf string `json:"cpf"`
}

func userHandler(response http.ResponseWriter, request *http.Request) {
	sid := strings.TrimPrefix(request.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
		case request.Method == "GET" && id > 0:
			userById(response, request, id)
		case request.Method == "GET":
			allUsers(response, request)
		default:
			response.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(response, "Sorry :(")
	}
}

func userById(response http.ResponseWriter, request *http.Request, id int) {
	database, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Não foi possível obter o usuário")
	}

	defer database.Close()

	var user User
	database.QueryRow("SELECT id, name, cpf FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Cpf)

	json, _ := json.Marshal(user)
	response.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(response, string(json))
}

func allUsers(response http.ResponseWriter, request *http.Request) {
	database, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Não foi possível obter os usuários")
	}

	defer database.Close()

	rows, _ := database.Query("SELECT id, name, cpf FROM users")
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Cpf)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)
	response.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(response, string(json))
}
