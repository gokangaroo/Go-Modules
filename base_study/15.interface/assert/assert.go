package main

import (
	"fmt"
)

func main() {
	var s1 interface{} = 56
	assert(s1)
	// 类型不一会报运行时异常
	var s2 interface{} = "56"
	assert(s2)
	// 为了能够适配所有类型, 可以使用v, ok := i.(T)语法
	// 添加一个ok来获取, 是否类型匹配, 类似于map取值, 如果没有就是0, 需要额外接受值
	findType("Naveen")
	findType(77)
	findType(89.98)
}
func assert(i interface{}) {
	s, ok := i.(int) //获取i变量底层int的值
	fmt.Println(s, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
