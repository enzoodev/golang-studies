package main

import (
	"fmt"
	"reflect"
)

func main() {
	firstArray := [3]int{1, 2, 3};
	firstSlice := []int{1, 2, 3};

	fmt.Println(reflect.TypeOf(firstArray), firstArray);
	fmt.Println(reflect.TypeOf(firstSlice), firstSlice);

	firstSlice = append(firstSlice, 7)

	// Slices aren't arrays, they are references to arrays

	// This is a slice of 2 elements of the firstArray
	secondSlice := firstArray[1:3]
	fmt.Println(secondSlice)

	// We can do a slice of a slice
	thirdSlice := secondSlice[0:1]
	fmt.Println(thirdSlice)
}
