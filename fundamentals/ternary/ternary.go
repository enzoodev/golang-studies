package main

// Go doesn't have the ternary operator

func getResult(note float64) string {
	if (note >= 7) {
		return "Approved"
	}

	return "Disapproved"
}

func main() {
	result := getResult(8.3)
	println(result)
}
