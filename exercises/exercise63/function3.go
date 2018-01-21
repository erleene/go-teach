package main

import "fmt"

func greatestNum(number ...int) int {
	num := 0
	for _, v := range number {
		if num < v {
			num = v
		}
	}
	return num
}

func main() {
	fmt.Println(greatestNum(500, 2, 3, 9, 10, 450))
}
