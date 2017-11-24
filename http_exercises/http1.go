package main

import (
	"fmt"
	"strings"
)

func worker(subworker func(s string)) {
	//do complex computation somewhere else...
	//pass to subworker

	subworker("super complex")
}
func main() {
	//two strings and return error
	var m map[string]func(s string, s1 string) error

	m = make(map[string]func(s string, s1 string) error)
	m["toy"] = func(s string, s1 string) error {
		fmt.Println(strings.Join([]string{s, s1}, " "))
		return nil
	}

	m["toy"]("gun", "ball")

	worker(
		func(s string) {
			fmt.Println(s)
		},
	)
}
