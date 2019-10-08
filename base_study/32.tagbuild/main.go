package main

import (
	"base_study/32.tagbuild/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

// go run -tags=third main.go
// 主要用来进行不同平台机器的不同方法使用 如 !windows 那么这个文件在windows上就不会被编译, 也不会调用里面的方法
func main() {
	u := user{"huija", 18}
	b, err := json.MarshalIndent(u, "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}
