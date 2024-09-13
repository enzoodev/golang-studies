package main

type Sportive interface {
	turnOnTurbo()
}

type Luxurious interface {
	doGoal()
}

type SportiveLuxurious interface {
	Sportive
	Luxurious
}

type BMW7 struct {}

func (b BMW7) turnOnTurbo() {
	println("Turbo is on")
}

func (b BMW7) doGoal() {
	println("Goal is done")
}

func main() {
	var b SportiveLuxurious = BMW7{}
	b.turnOnTurbo()
	b.doGoal()
}
