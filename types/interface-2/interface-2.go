package main

type Sportive interface {
	turnOrTurbo()
}

type Ferrari struct {
	model string
	speed int
	isTurboOn bool
}

func (f *Ferrari) turnOrTurbo() {
	f.isTurboOn = true
}

func main() {
	firstCar := &Ferrari{
		model: "Ferrari 488 GTB",
		speed: 330,
		isTurboOn: false,
	}
	firstCar.turnOrTurbo()

	secondCar := &Ferrari{
		model: "Ferrari 812 Superfast",
		speed: 340,
		isTurboOn: false,
	}
	secondCar.turnOrTurbo()

	println(firstCar.model, "Turbo is on:", firstCar.isTurboOn)
	println(secondCar.model, "Turbo is on:", secondCar.isTurboOn)
}
