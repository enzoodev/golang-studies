package main

func getSum(a, b int) int {
	return a + b
}

func getSumAndDiff(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	myFirstSum := getSum(2, 3)
	mySecondSum, myFirstDiff := getSumAndDiff(2, 3)

	println(myFirstSum)
	println(mySecondSum, myFirstDiff)
}
