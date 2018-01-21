package main

import "fmt"

//swap two ints (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	fmt.Println(x, y) //should print x=3 an y=2
	swap(&x, &y)
	fmt.Println(x, y)
}
