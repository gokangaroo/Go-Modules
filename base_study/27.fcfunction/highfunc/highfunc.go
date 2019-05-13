package main

import (
	"fmt"
)

//1. 函数作为参数
func simple1(c func(a, b int) int) {
	fmt.Println(c(60, 7))
}

//2. 函数作为返回值
func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	// 返回定义的函数
	return f
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple1(f)

	// s是返回的函数, 类型为func(a, b int)
	s := simple2()
	// 这个参数有点难懂
	fmt.Println(s(60, 7))
}
