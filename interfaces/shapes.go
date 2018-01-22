package main

type Shape struct {
	name string
}

type Triangle struct {
	Shape
}
type Square struct {
	Shape
}
type Circle struct {
	Shape
}

type IShape interface {
	getArea() float64	name string

	getName() string
}

func (t Triangle) getArea() float64 {
	//implementation
	return 0
}
func (t Triangle) getName() string {
	return t.name
}

func (c Circle) getArea() float64 {
	//implementation
	return 0
}

func (s Square) getArea() float64 {
	//implementation
	return 0
}

func GetArea(i IShape) float64 {
	return i.getArea()
}

func Getname(i IShape) string {
	return i.getName()
}

func main() {
	triangle1 := &Triangle{Shape: Shape{name: "Triangle"}}

	GetArea(triangle1)
	GetName(triangle1)

	square1 := new(Square)

	GetArea(square1)
	GetName(square1)

	circle1 := new(Circle)

	GetArea(circle1)
	GetName(square1)
}
