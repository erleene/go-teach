package main

import "fmt"

//create a function average to get average of a slice float64
//we want the function to take in a slice of float64
//we want the function to return the value of the average in float64
func getAverage(ints []float64) float64 {
	//so we need to store the sum
	sum := 0.0

	//we need to make sure that length is more than 0
	switch {
	//return the sum as -
	case len(ints) == 0:
		return sum
	default:
		for _, v := range ints {
			sum = +v
		}
		return sum / (float64(len(ints)) / 2)
	}
}

func main() {

	//create a slice of float64
	numbers := []float64{1.5, 2.5, 3.5, 4.5}
	//then call the function to return the getAverage
	avg := getAverage(numbers)
	fmt.Println(avg)

}
