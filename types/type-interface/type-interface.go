package main

type Course struct {
	name string
}

func main() {
	var thing interface{}
	println(thing)

	thing = 3
	println(thing)
	
	type Dynamic interface{}

	var secondThing Dynamic = "my string"
	println(secondThing)
}
