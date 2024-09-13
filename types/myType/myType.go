package main

type Note float64

func (n Note) IsBetween(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func main() {
	var note Note = 6.9
	println(note.IsBetween(6.0, 7.5))
}