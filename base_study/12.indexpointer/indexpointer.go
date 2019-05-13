package main

import (
	"fmt"
)

func main() {
	b := 255
	// 跟c语言一样的指针声明, 指针默认是nil
	var a *int
	fmt.Println("index a is", a)
	a = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	// 指针的解引用时*a, 可以获取到值
	fmt.Println("value of b is", *a)
	// 同时解引用可以直接操作内存
	fmt.Println("value of b+1 is", *a+1)
	// 函数传递指针参数
	change(a)
	fmt.Println("after function ,value of b is", *a)
	// 1. 传递数组指针, 修改数组
	c := [3]int{101, 102, 91}
	modify(&c) // 指向整个数组的指针, 类型为 *[3]int
	fmt.Println(c)
	// 2.最好使用切片, 而不是数组指针
	d := [3]int{101, 102, 91}
	modifyslice(d[:])
	fmt.Println(d)
	// 不支持指针运算如 a++从b[0]指向b1
	// e := [...]int{109, 110, 111}
	// p := &e[0]
	// p++
}

// 修改内存值为55
func change(val *int) {
	*val = 55
}

// 根据数组指针修改数组
func modify(arr *[3]int) {
	(*arr)[0] = 89
	//a[x] 是 (*a)[x] 的简写形式
	arr[1] = 90
}

// 根据切片来修改数组
func modifyslice(sls []int) {
	sls[0] = 89
	sls[1] = 90
}
