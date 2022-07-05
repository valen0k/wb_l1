package ex24

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func Distance(first, second Point) float64 {
	return math.Sqrt(math.Pow(first.x-second.x, 2) + math.Pow(first.y-second.y, 2))
}
