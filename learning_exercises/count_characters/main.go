package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "asSaSA ddd dsjkdsjs dk"
	fmt.Println(len([]rune(text)))
	//print the byte
	rune_text := []rune(text)
	for _, v := range rune_text {
		fmt.Println(utf8.RuneLen(v))
	}
}
