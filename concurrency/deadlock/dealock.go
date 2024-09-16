package main

import "time"

func routine(c chan int) {
	time.Sleep(1 * time.Second)
	c <- 15
	println("routine done")
}

func main() {
	c := make(chan int) // unbuffered channel

	go routine(c)

	println(<-c)
	// println("main done")
	// println(<-c)
}
