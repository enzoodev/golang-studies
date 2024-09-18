package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func correctTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("06/05/2003 03:04:05")
	fmt.Fprintf(w, "<h1>The correct time is: %s<h1>", s)
}

func main() {
	http.HandleFunc("/correctTime", correctTime)
	const port = "8080"

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
