package main

import (
	"fmt"
)

func main() {
	// 简单操作
	find(89, 89, 90, 95)
	find(87)
	// 传入数组
	nums := []int{89, 90, 95}
	//find(89, nums) //无法传入, []{nums}
	find(89, nums...) //解决方法就是加...
	// 引用传递
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

// V1.0可变参数函数
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums) // %T获取参数类型
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

// V2.0函数是引用传递
func change(s ...string) {
	s[0] = "Go"                 // 这个是生效的, 因为是直接操作地址.
	s = append(s, "playground") //这个出去就失效了,因为append会进行值传递, 创建切片副本.
	fmt.Println(s)
}
