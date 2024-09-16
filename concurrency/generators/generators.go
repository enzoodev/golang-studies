package main

import (
	"io"
	"net/http"
	"regexp"
)

// <-channel is a receive expression

func getTitle(urls ...string) <-chan string {
	channel := make(chan string)

	for _, url := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := io.ReadAll(response.Body)

			r, _ := regexp.Compile("<title>(.*)</title>")
			channel <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return channel
}

func main() {
	firstTitle := getTitle("https://www.google.com", "https://www.facebook.com")
	secondTitle := getTitle("https://www.instagram.com", "https://www.youtube.com")

	println("firsters: ", <-firstTitle, <-secondTitle)
	println("seconders: ", <-firstTitle, <-secondTitle)
}
