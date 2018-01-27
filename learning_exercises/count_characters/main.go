package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "asSaSA ddd dsjkdsjs dk"
	text_rune := []rune(text)
	copy(text_rune[4:4+3], []rune("abc"))

	fmt.Println(len([]byte(text)))
	fmt.Println(utf8.RuneCount([]byte(text)))
	//print the byte
	// rune_text := []rune(text)
	// for _, v := range rune_text {
	//
	// 	fmt.Println(v)

}
