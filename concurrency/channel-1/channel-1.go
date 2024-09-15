package main

func main() {
	channel := make(chan int, 1)

	channel <- 1 // send value 1 to the channel
	<-channel    // receive value from the channel

	channel <- 2 // send value 2 to the channel
	println(<-channel) // receive value from the channel and print it
}
