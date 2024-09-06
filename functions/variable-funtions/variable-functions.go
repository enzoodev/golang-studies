package main

import "fmt"

func media(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("%.2f \n", media(7.7, 8.2, 9.1, 10.7))
	fmt.Printf("%.2f", media(7.7, 8.2, 9.1))
}
