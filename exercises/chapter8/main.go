package main

import (
	"fmt"

	"github.com/go-teach/exercises/chapter8/math"
)

func main() {
	xs := []float64{1, 2, 3, 4, 9, 10}
	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println("average: ", avg)
	fmt.Println("min:", min)
	fmt.Println("max:", max)
}
