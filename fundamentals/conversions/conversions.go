package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	note := 6.9
	finalNote := int(note)
	fmt.Println(finalNote)

	// Careful...
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("Test " + strconv.Itoa(97))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 102)
	
	// string to boolean
	b, _ := strconv.ParseBool("1")

	fmt.Println(b)
	if b {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}