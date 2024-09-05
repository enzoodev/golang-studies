package main

import "fmt"

func getConcept(note float64) (string) {
	roundedNote := int(note)

	switch {
		case roundedNote >= 9:
			return "A"
		case roundedNote >= 7:
			return "B"
		case roundedNote >= 5:
			return "C"
		case roundedNote >= 3:
			return "D"
		case roundedNote >= 0:
			return "E"
		default:
			return "Invalid note"
	}
}

func main() {
	fmt.Println(getConcept(9.8))
	fmt.Println(getConcept(6.9))
}
