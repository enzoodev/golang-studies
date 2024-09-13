package main

func factorial(n uint) uint {
	switch {
		case n == 0:
			return 1
		default:
			return n * factorial(n-1)
	}
}

func main() {
	fatorialFour  := factorial(4)
	println(fatorialFour)

	fatorialZero  := factorial(0)
	println(fatorialZero)
}
