package main

type Car struct {
	name string
	currentSpeed int
}

type Ferrari struct {
	Car
	price int
}

func main() {
	ferrari := Ferrari{}
	ferrari.name = "Ferrari"
	ferrari.currentSpeed = 0
	ferrari.price = 1000000
	
	
	println(ferrari.name)
	println(ferrari.currentSpeed)
	println(ferrari.price)
}