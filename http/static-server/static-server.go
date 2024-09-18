package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/", fileServer)

	const port = "8080"

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil)) 
}
