package main

import "strconv"

type Printer interface {
	toString() string
}

type Person struct {
	name string
	surname string
}

type Product struct {
	name string
	price float64
}

// interfaces are implemented implicitly

func (p Person) toString() string {
	return p.name + " " + p.surname
}

func (p Product) toString() string {
	return p.name + " - $ " + strconv.FormatFloat(p.price, 'E', -1, 32)
}

func print(p Printer) {
	println(p.toString())
}

func main() {
	person := Person{"John", "Doe"}
	product := Product{"Laptop", 2500.00}

	print(person)
	print(product)
}
