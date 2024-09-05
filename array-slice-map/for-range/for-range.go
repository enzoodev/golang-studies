package main

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

	for index, number := range numbers {
		println(index, number)
	}

	for _, number := range numbers {
		println(number)
	}
}
