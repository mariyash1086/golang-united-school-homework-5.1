package square

//package main

type Point struct {
	x, y int
}
type Square struct {
	start Point
	a     uint
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

	resultPoint.x = int(c.a) + c.start.x
	resultPoint.y = int(c.a) + c.start.y
	return resultPoint
}

func (c Square) Area() uint {

	return c.a * c.a

}

func (c Square) Perimeter() uint {
	return c.a * 4
}
