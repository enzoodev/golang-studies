package main

func change(p1, p2 int) (second int, first int) {
	first = p1
	second = p2
	
	return
}

func main() {
	first, second := change(10, 20)
	println(first, second)
}
