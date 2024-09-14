package main

import "math"

// initiliazing with a capital letter to make it public
// initiliazing with a small letter to make it private

type Point struct {
	x, y float64
}

func calculatePeccaries(p1, p2 Point) (px, py float64) {
	px = math.Abs(p1.x - p2.x)
	py = math.Abs(p1.y - p2.y)

	return px, py
}

func CalculateDistance(p1, p2 Point) float64 {
	px, py := calculatePeccaries(p1, p2)
	return math.Sqrt(math.Pow(px, 2) + math.Pow(py, 2))
}
