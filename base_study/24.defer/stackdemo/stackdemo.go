package main

import (
	"fmt"
)

func main() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	//遍历字符串, 逆序打印字符
	for _, v := range []rune(name) {
		// 10个字符就10个defer
		defer fmt.Printf("%c", v)
	}
}
