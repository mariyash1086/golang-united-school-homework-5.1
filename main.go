package golang_united_school_homework_5_1

//package main

import (
	"math"
)

type Point struct {
	x, y int
}
type Square struct {
	start Point
	side  uint
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

	resultPoint.x = int(c.side) + c.start.x
	resultPoint.y = int(c.side) + c.start.y
	return resultPoint
}

func (c Square) Area() float64 {

	return math.Pow(float64(c.side), 2)

}

func (c Square) Perimeter() float64 {
	return float64(c.side) * 4
}
