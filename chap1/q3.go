package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tree()
	count("asSASA ddd dsjkdsjs dk")
}
func tree() {
	for i := 0; i < 100; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("A")
		}
	}
}
func count(str string) int {
	return utf8.RuneCountInString(str)
}
