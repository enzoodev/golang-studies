package main

import "fmt"

func main() {
	slice := make([]int, 5)
	slice[4] = 19
	fmt.Println(slice)

	slice = make([]int, 5, 5)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5, 6)
	fmt.Println(slice, len(slice), cap(slice))
}
