package main

import "github.com/enzoodev/formatHtml"

func forward(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin
	}

}

func multiplexer(origin1, origin2 <-chan string) <-chan string {
	channel := make(chan string)

	go forward(origin1, channel)
	go forward(origin2, channel)

	return channel
}

func main() {
	channel := multiplexer(
		formatHtml.GetTitle("https://www.google.com", "https://www.facebook.com"),
		formatHtml.GetTitle("https://www.google.com", "https://www.facebook.com"),
	)

	println(<-channel)
}
