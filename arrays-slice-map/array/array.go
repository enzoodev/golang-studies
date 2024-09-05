package main

import "fmt"

func main() {
	// Arrays in go are static, meaning that once you define the size of an array, you can't change it.

	// This is an array of 5 integers
	var myArray [5]int = [5]int{1, 2, 3, 4, 5}
	myArray[0], myArray[3] = 5, 9

	total := 0.0
	for i := 0; i < len(myArray); i++ {
		total += float64(myArray[i])
	}

	median := total / float64(len(myArray))

	fmt.Println(median)
	fmt.Println(myArray)
}
