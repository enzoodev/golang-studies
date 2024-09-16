package main

import "time"

func isCousin(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func cousins(n int, channel chan int) {
	start := 2

	for i := 0; i < n; i++ {
		for cousin := start; ; cousin++ {
			if isCousin(cousin) {
				channel <- cousin
				start = cousin + 1
				time.Sleep(100 * time.Millisecond)
				break
			}
		}
	}

	close(channel)
}

func main() {
	channel := make(chan int, 30)
	go cousins(cap(channel), channel)

	for cousin := range channel {
		println(cousin)
	}
}
