package main

func getApprovedValue(value int) int {
	defer println("End of getApprovedValue")

	if value > 10 {
		println("Value is greater than 10")
		return value
	}

	println("Value is less than 10")
	return 10
}

func main() {
	println(getApprovedValue(20))
}
