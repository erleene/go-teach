package main

import "fmt"

var smallest int

func main() {

	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}

	//find the smallest number in the array x
	// //do this by comparing each if the lenght is zero get the first one
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
		fmt.Println(x[i+1])
		// /if x[i] < x[i+1] {
		// 	//store the small one
		// 	smallest = x[i]
		// }
	}
	fmt.Println(len(x))

}
