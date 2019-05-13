package main

import (
	"fmt"
)

// 1.函数也是变量,你怕不怕, 这种没有名字的函数叫做匿名函数
var a = func() {
	fmt.Println("hello world first class function")
}

// 2.那么可不可以自定义一个类型, 然后把函数给这个类型呢?
// 定义add类型
type add func(a int, b int) int

func main() {
	a()
	fmt.Printf("%T", a)
	func(n string) {
		fmt.Println("匿名函数 ", n)
	}("你怕不怕")
	// 你怕不怕
	var b add = func(a int, b int) int {
		return a + b
	}
	s := b(5, 6)
	fmt.Println("Sum", s)
}
