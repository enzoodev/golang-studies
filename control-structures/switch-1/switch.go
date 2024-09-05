package main

func noteForConcecpt(note float64) (string) {
	roundedNote := int(note)

	switch roundedNote {
		case 10: 
			fallthrough
		case 9:
			return "A"
		case 8, 7:
			return "B"
		case 6, 5:
			return "C"
		case 4, 3:
			return "D"
		case 2, 1, 0:
			return "E"
		default:
			return "Invalid note"
	}
}

func main() {
	println(noteForConcecpt(9.8))
	println(noteForConcecpt(6.9))
	println(noteForConcecpt(2.3))
}
