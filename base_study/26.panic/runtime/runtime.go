package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	a()
	fmt.Println("normally returned from main")
}

// recover函数
func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}

// panic函数
func a() {
	defer r()
	n := []int{5, 7, 4}
	// 数组越界
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}
