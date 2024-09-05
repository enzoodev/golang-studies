package main

import "fmt"

func main() {
	firstArray := [3]int{1, 2, 3}
	var firstSlice []int

	firstSlice = append(firstSlice, 4, 5, 6)
	fmt.Println(firstArray, firstSlice)

	secondSlice := make([]int, 2)

	fmt.Println(firstSlice, secondSlice)

	copy(secondSlice, firstSlice)

	fmt.Println(firstSlice, secondSlice)
}
