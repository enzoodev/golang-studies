package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	j := 0
	for j < 10 {
		println(j)
		j++

		if j % 2 == 0 {
			println("Even")
		} else {
			println("Odd")
		}
	}

	// infinite loop

	for {
		println("Infinite loop")
		time.Sleep(time.Second)
	}
}