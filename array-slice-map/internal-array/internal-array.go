package main

import "fmt"

func main() {
	firstSlice := make([]int, 10, 20)
	secondSlice := append(firstSlice, 1, 2, 3, 4, 5)

	fmt.Println(firstSlice, len(firstSlice), cap(firstSlice))
	fmt.Println(secondSlice, len(secondSlice), cap(secondSlice))

	firstSlice[0] = 42

	// firstSlice and secondSlice share the same underlying array
	fmt.Println(firstSlice, len(firstSlice), cap(firstSlice))
	fmt.Println(secondSlice, len(secondSlice), cap(secondSlice))
}
