package main

import "math"

func main() {
	const PI float64 = 3.14
	var radius = 3.3

	area := PI * math.Pow(radius, 2)

	println("Circle area is", area)

	const (
		age int = 21
	)

	johnName, enzoName := "John", "Enzo"

	enzoName = "Enzo tt"

	println("My age is", age)
	println("My name is", enzoName)
	println("My name is", johnName)
}
