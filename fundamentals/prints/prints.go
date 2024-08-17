package main

import "fmt"

func main() {
	fmt.Print("Hello")
	fmt.Println(" Line")
	aleatoryNumber := 10.1236165151
	fmt.Printf("\nThe number is %.2f", aleatoryNumber)

	aleatoryNumberFormattedToString := fmt.Sprint(aleatoryNumber)

	fmt.Println("\nThe number is " + aleatoryNumberFormattedToString)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
