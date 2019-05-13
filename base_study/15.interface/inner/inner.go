package main

import (
	"fmt"
)

// Test :接口
type Test interface {
	Tester()
}

// MyFloat :自定义类型来实现接口
type MyFloat float64

// Tester :打印自身
func (m MyFloat) Tester() {
	fmt.Println(m)
}

// 输出接口的实际类型和值: 跟java其实是一样的, 接口是编译类型, 具体接收器是实际类型
func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}
