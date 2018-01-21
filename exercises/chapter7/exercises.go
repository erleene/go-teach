package main

//a method is used on a struct
//a function can be used anywhere
//embedded anonymous field is-a relationship with a struct
type Circle struct {
}

type Rectangle struct {
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (c *Circle) perimeter() float64 {

}

func (r *Rectangle) perimeter() float64 {

}

func main() {

}
