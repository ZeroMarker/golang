package main

import (
	"fmt"
	"gopher/chap1"
)

func main() {
	fmt.Println("Starting...")
	insertTime()
	fmt.Println("return", test())
	fmt.Println("return", other())
	chap1.Hello()
}

// defer run after return, before function end.
func test() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

// return variable has name, defer function will change the returned value.
func other() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer")
	}()
	return i
}
