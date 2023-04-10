package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func lenght(p1 Point, p2 Point) float64 { //более наглядно, можно было через math.pow
	a := p2.x - p1.x
	b := p2.y - p1.y
	res := a*a + b*b
	result := math.Sqrt(float64(res))
	return result
}

func main() {
	p1 := Point{4, 2}
	p2 := Point{1, 6}

	fmt.Printf("result is: %v", lenght(p1, p2))
}
