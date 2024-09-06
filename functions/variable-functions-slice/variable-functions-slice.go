package main

func printApproved(approved ...string) {
	for i, student := range approved {
		println(i+1, student)
	}
}

func main() {
	approved := []string{"Mary", "Peter", "John", "Alice"}
	printApproved(approved...)
}
