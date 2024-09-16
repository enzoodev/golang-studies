package main

import (
	"fmt"
	"time"
)

func talk(person string) <- chan string {
	channel := make(chan string)

	go func(){
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			channel <- fmt.Sprintf("%s says: %d", person, i)
		}
	}()

	return channel
}

func multiplexer(firstEntry, secondEntry <- chan string) <- chan string {
	channel := make(chan string)

	go func(){
		for {
			select {
				case message := <- firstEntry:
					channel <- message
				case message := <- secondEntry:
					channel <- message
			}
		}
	}()

	return channel
}

func main() {
	channel := multiplexer(talk("Alice"), talk("Bob"))

	for i := 0; i < 6; i++ {
		fmt.Println(<-channel)
	}
}
