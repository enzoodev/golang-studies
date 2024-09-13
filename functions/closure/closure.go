package main

func closure() func() {
	x := 10
	
	var f = func() {
		println(x)
	}

	return f
}

func main() {
	x := 20
	println(x)

	printX := closure()
	printX()
}
