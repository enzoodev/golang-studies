package main

import "time"

func twoThreeFour(base int, c chan int) {
	time.Sleep(1 * time.Second)
	c <- 2 * base

	time.Sleep(1 * time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	channel := make(chan int)
	go twoThreeFour(2, channel)
println("A")
	a, b := <-channel, <-channel
println("B")
	println(a, b)
	println(<-channel)
}
