package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates:", d1.Equal(d2))

	type Person struct {
		name string
		age int
	}

	p1 := Person{"John", 12}
	p2 := Person{"John", 12}

	fmt.Println("Persons:", p1 == p2)
}