package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
		case n < 0:
			return 0, fmt.Errorf("invalid number: %d", n)
		case n == 0:
			return 1, nil
		default:
			value, _ := factorial(n-1)
			return n * value, nil
	}
}

func main() {
	fatorialFour, _  := factorial(4)
	println(fatorialFour)

	fatorialZero, _  := factorial(0)
	println(fatorialZero)

	_, err := factorial(-4)
	if err != nil {
		println(err)
	}
}
