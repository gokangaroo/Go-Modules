package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()

	fmt.Println(a("World"))
	fmt.Println(a("Everyone"))
	// 每一个闭包,会绑定外围变量, 所以这边会接在后面
	fmt.Println(a("Gopher"))
	fmt.Println(a("!"))
}
