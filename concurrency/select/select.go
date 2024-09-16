package main

import (
	"io"
	"net/http"
	"regexp"
	"time"
)

func GetTitle(urls ...string) <-chan string {
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

func getMostFastest(url1, url2, url3 string) string {
	firstChannel := GetTitle(url1)
	secondChannel := GetTitle(url2)
	thirdChannel := GetTitle(url3)

	select {
		case firstTitle := <-firstChannel:
			return firstTitle
		case secondTitle := <-secondChannel:
			return secondTitle
		case thirdTitle := <-thirdChannel:
			return thirdTitle
		case <-time.After(time.Millisecond * 1000):
			return "All requests are slow"
		default:
			return "No response"
	}
}

func main() {
	champion := getMostFastest(
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	)

	println(champion)
}
