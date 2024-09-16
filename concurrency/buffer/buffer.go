package main

import "time"

func routine(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	println("routine execute")
	channel <- 5
	channel <- 6
}

func main() {
	channel := make(chan int, 3) // buffered channel

	go routine(channel)

	time.Sleep(time.Second)
	println(<-channel)
	// println(<-channel)
	// println(<-channel)
	// println(<-channel)
	// println(<-channel)
	// println(<-channel)
}
