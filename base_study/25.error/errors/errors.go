package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func main() {
	// 异常直接比较,有些返回的是固定的异常类型
	// 底层其实就是: var ErrBadPattern = errors.New("syntax error in pattern")
	files, error := filepath.Glob("/")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	// 简单自定义一个异常, 另外建一个circle.go来实战一下
	fmt.Println("matched files", files)
	var MyselfError = errors.New("自定义异常")
	fmt.Println(MyselfError)
}
