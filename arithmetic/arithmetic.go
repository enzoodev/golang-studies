package main

import "math"

func main() {
	a:= 3
	b:= 2

	// arithmetic operations

	// sum
	sum := a + b
	println("sum", sum)

	// subtraction
	sub := a - b
	println("sub", sub)

	// multiplication
	mult := a * b
	println("mult", mult)

	// division
	div := a / b
	println("div", div)

	// module
	mod := a % b
	println("mod", mod)

	// bitwise operations
	binay := a&b
	println("binary", binay)

	binay = a|b
	println("binary", binay)

	binay = a^b
	println("binary", binay)

	// Power
	pow := math.Max(float64(a), float64(b))
	println("pow", pow)

	// shift operations
	shift := a << 1
	println("shift", shift)

	shift = a >> 1
	println("shift", shift)
}