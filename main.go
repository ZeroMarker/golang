package main

import (
	"fmt"
	"gopher/chap1"
	"github.com/ZeroMarker/cnid"
)

func main() {
	fmt.Println("Starting...")
	//insertTime()
	fmt.Println("return", test())
	fmt.Println("return", other())
	chap1.Hello()

	cases := []string{
		"11010519491231002X", // 正确
		"110105194912310021", // 校验码错误
		"11010519490230002X", // 日期错误
	}

	for _, c := range cases {
		if err := cnid.ValidateIDCard(c); err != nil {
			println(c, "->", err.Error())
		} else {
			println(c, "-> OK")
		}
	}
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
