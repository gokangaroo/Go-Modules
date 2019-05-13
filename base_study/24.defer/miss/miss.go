package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("那么defer函数使用的值为", a)
}
func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("参数修改之后的值为", a)

}
