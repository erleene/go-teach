package main

import "fmt"

func half(number int) (int, bool) {
	half_of_num := number / 2

	//check if the half is even
	if number%2 == 0 {
		return half_of_num, true
	}
	return half_of_num, false
}

func main() {
	nums := []int{1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println(half(nums[i]))
	}
}
