package main

import "fmt"

func main() {
	// var approved map[string]string
	// maps must be initialized
	approved := make(map[string]string)

	approved["07929725545"] = "Enzo"
	approved["07929725546"] = "Maria"
	approved["07929725547"] = "João"
	
	fmt.Println(approved)

	for key, value := range approved {
		fmt.Printf("%s: %s\n", key, value)
	}

	fmt.Println(approved["07929725545"])
	delete(approved, "07929725546")
	fmt.Println(approved)
}
