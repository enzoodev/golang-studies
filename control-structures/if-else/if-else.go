package main

func printResult(note float64) {
	if note >= 7 {
		println("Result is true")
	} else {
		println("Result is false")
	}
}

func main() {
	printResult(9.5)
	printResult(4)
}