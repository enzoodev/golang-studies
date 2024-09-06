package main

var sum = func(a, b int) int {
	return a + b
}

func main() {
	println(sum(2, 3))

	diff := func(a, b int) int {
		return a - b
	}

	println(diff(2, 3))
}
