package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// int numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("Litera integer is", reflect.TypeOf(32000))

	// unsigned int numbers
	// uint8 uint16 uint32 uint64

	const b byte = 255
	fmt.Println("The byte is", reflect.TypeOf(b))

	maxInt8 := math.MaxInt8
	fmt.Println("The max value of int8 is", maxInt8)

	maxInt16 := math.MaxInt16
	fmt.Println("The max value of int16 is", maxInt16)

	maxInt32 := math.MaxInt32
	fmt.Println("The max value of int32 is", maxInt32)

	maxInt64 := math.MaxInt64
	fmt.Println("The max value of int64 is", maxInt64)

	const myRune rune = 'a' // represents a unicode table (int32)
	fmt.Println("The rune is", reflect.TypeOf(myRune))
	fmt.Println(myRune)

	// real numbers (float32, float64)
	const x float32 = 49.99
	fmt.Println("The float is", reflect.TypeOf(x))

	// boolean
	boo := true
	fmt.Println("The boolean is", reflect.TypeOf(boo))

	// string
	s1 := "Hello, my name is Enzo"
	fmt.Println(s1 + "!")
	fmt.Println("The string is", reflect.TypeOf(s1))
	fmt.Println(("The string length is"), len(s1))

	// string with multiple lines
	s2 := `Hello
	my
	name
	is
	Enzo`

	fmt.Println(s2)

	// char (not exist in Go)
	// var x char = 'b'
	char := 'a'
	fmt.Println("The char is", reflect.TypeOf(char))
}
