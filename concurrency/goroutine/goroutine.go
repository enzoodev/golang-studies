package main

import "time"

func talk(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(1 * time.Second)
		println(person, "says:", text, "quantity:", i)
	}
}

func main() {
	go talk("Alice", "Hello", 50)
	go talk("Bob", "Hi", 50)

	// if we don't wait for the goroutines to finish, the program will exit
	println("Waiting for the talks to finish...")
}