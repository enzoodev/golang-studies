package main

func noteForConcecpt(note float64) {
	if note >= 9 {
		println("A")
	} else if note >= 8 {
		println("B")
	} else if note >= 7 {
		println("C")
	} else if note >= 6 {
		println("D")
	} else {
		println("F")
	}
}

func main() {
	noteForConcecpt(9.8)
	noteForConcecpt(6.9)
	noteForConcecpt(2.3)
}
