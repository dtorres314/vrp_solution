package main

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) distanceTo(other Point) float64 {
	return math.Sqrt((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))
}

type Load struct {
	id      string
	pickup  Point
	dropoff Point
}
