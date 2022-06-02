package golang_united_school_homework_5_1

//package main

import (
	"math"
)

type Point struct {
	x, y int32
}
type Square struct {
	Start Point
	A     int32
}

//func main() {

//	example := Square{Y: 5}
//	a := example.perimeter(example)
//	b := example.endPoint(example)
//	c := example.squareMetod(example)

//	fmt.Println(a, b, c)

//}

func (c Square) End() Point {

	var resultPoint Point

	resultPoint.x = c.A + c.Start.x
	resultPoint.y = c.A + c.Start.y
	return resultPoint
}

func (c Square) Area() float64 {

	return math.Pow(float64(c.A), 2)

}

func (c Square) Perimeter() float64 {
	return float64(c.A) * 4
}
