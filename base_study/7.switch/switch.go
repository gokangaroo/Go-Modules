package main

import (
	"fmt"
)

func main() {
	switch letter := "i"; letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	num := 75
	switch { // 缺省匹配true的case,没有fallthrough就匹配第一个true的case
	case num > 0:
		fmt.Println("num is greater than 0")
		fallthrough
	case num < 50:
		fmt.Println("num is greater than 0 and less than 50")
		fallthrough //fallthrough表示无条件运行下一个case,无视是否匹配
	case num >= 101:
		fmt.Println("num is greater than 100")
		fallthrough
	case num >= 51:
		fmt.Println("num is greater than 51")
	}
}
