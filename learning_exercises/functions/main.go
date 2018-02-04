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

//Write a function that returns its (two)parameters in the right
// numerical(ascend- ing) order:
//  f(7,2) → 2,7     so it checks to see if the current value is less then the next one, if so it prints it,
//   f(2,7) → 2,7
func numericalAscending(x, y int) (int, int) {
	//check if x is less than y
	switch {
	case x < y:
		return x, y
	default:
		return y, x
	}
}

//Create a simple stack which can hold a fixed number of ints.
//It does not have to grow beyond this limit.
type stack struct {
	i    int
	data [10]int //can store dat. data[0] =  2, data[1] = 3. i will be the index
}

//Define push – put something on the stack
//push will push to the bottom of the stack stack(len()-1)
func (s *stack) push(x int) {
	//if stack is empty
	//we want to put x on the bottom of the stack
	s.data[s.i] = x
	s.i++
}

// func (s *stack) pop() []int {
// 	//if not empty
// 	//remove x from the top of the stack
// 	s.i--
// 	//remove from data from the top
// 	top := 0
// 	//element to remove
// 	//we need to create a new data int[] to return
//
// 	return s.data[top++:i]
//
// }

//Write a function that takes a variable number of ints and prints each integer on a separate line
func printInt(num ...int) {
	for _, v := range num {
		fmt.Println("%d\n", v)
	}
}

//Map function A map()-function is a function
//that takes a function and a list.
//The function is applied to each member in the list
//and a new list containing these calculated values is returned. Thus:
//map(f(), (a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))
func map(f func(), ...int) ...func() {
	
}



//The Fibonacci sequence starts as follows:1,1,2,3,5,8,13,... Or in mathematical
//terms: x1 = 1;x2 = 1;xn = xn−1 +xn−2 ∀n > 2.
//Write a function that takes an int value and gives that many terms of the Fibonacci
//sequence.
func fibonacci(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

func main() {
	var s stack
	for _, c := range fibonacci(10) {
		fmt.Printf("%v", c)
	}

	//create a slice of float64
	numbers := []float64{1.5, 2.5, 3.5, 4.5}
	//then call the function to return the getAverage
	avg := getAverage(numbers)
	fmt.Println(avg)
	fmt.Println(numericalAscending(7, 2))
	fmt.Println(numericalAscending(2, 7))
	fmt.Println(s)
	s.push(45)
	s.push(46)
	fmt.Println(s)
	// s.pop()
	fmt.Println(s)

	printInt(1, 2, 3, 4)

}
