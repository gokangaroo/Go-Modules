package main

import (
	"fmt"
)

// 空接口,可以接受所有类型, 因为其没有什么方法需要实现
type myinter interface {
}

func describe(i myinter) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}
