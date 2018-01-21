package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var s, sed string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += sed + os.Args[i]
	// 	sed = " "
	// }
	// fmt.Println(s)
	// fmt.Println(len(os.Args))
	s, sed := "", ""
	for _, arg := range os.Args[1:] {
		s += sed + arg
		sed = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
