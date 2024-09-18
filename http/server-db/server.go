package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", userHandler)

	const port = "8080"

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
