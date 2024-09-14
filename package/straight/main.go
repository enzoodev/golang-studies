package main

import "fmt"

func main() {
	p1 := Point{1, 2}
	p2 := Point{4, 6}
	distance := CalculateDistance(p1, p2)

	fmt.Println(distance)
}