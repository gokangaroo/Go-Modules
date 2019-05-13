package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

// 使用recover对panic进行处理
func recoverName() {
	// 如果捕获到panic,跳过异常语句
	if r := recover(); r != nil {
		fmt.Println("处理开始,recovered from ", r)
	}
}

// 打印一个人的全名
func fullName(firstName *string, lastName *string) {
	defer recoverName()
	//defer fmt.Println("deferred call in fullName")
	// 这里如果不加nil判断,就会panic: runtime error: invalid memory address or nil pointer dereference
	// 所以需要panic上场
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
