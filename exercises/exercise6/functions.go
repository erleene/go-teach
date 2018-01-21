package main

import "fmt"

func sum(slice []int) int {
	//
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func main() {

	//create an array of numb
	num_arrays := []int{48, 46, 47, 67, 89}

	//get the slice of numbers:48 46
	slice_of_num := num_arrays[0:2]
	//add the numbers
	total_num := sum(slice_of_num)

	fmt.Println(total_num)

}
