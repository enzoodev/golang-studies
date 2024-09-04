package main

import (
	"math/rand"
	"time"
)

func aleatoryNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i := aleatoryNumber(); i > 5 {
		println("Win")
	} else if i > 3 {
		println("Lose a little")
	} else {
		println("Lose a lot")
	}
}
