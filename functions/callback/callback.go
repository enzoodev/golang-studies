package main

func multiply(a, b int) int {
	return a * b
}

func execute(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main() {
	multiplyResult := execute(multiply, 2, 3)
	println(multiplyResult)
}
