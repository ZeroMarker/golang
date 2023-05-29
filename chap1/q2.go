package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 { // first judge
			fmt.Println("Fizz")
		} else {
			if i%5 == 0 { // second judge
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}
	}
}
