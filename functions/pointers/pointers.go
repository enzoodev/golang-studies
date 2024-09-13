package main

func incrementValueWithoutPointer(value int) {
	value++
}

func incrementValueWithPointer(value *int) {
	*value++
}

func main() {
	value := 10
	println("Value before increment: ", value)
	incrementValueWithoutPointer(value)
	println("Value after increment without pointer: ", value)
	incrementValueWithPointer(&value)
	println("Value after increment with pointer: ", value)

	// & is the address of operator
}
