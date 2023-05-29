package main

import "fmt"

func main() {
	forLoop()
	gotoLoop()
	gotoArrayLoop()
}

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func gotoLoop() {
	i := 0
	goto Here
	// goto label
Here:
	if i < 10 {
		fmt.Println(i)
		i++
		goto Here
	} else {

	}

}
func gotoArrayLoop() {
	age := [10]int{}
	for i := 0; i < 10; i++ {
		age[i] = i
	}
	i := 0
	goto There
	// goto label
There:
	if i < 10 {
		fmt.Println(age[i])
		i++
		goto There
	} else {

	}
}
