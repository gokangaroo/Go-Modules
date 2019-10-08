package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 1.make创建
	personSalary1 := make(map[string]int)
	personSalary1["steve"] = 12000
	personSalary1["jamie"] = 15000
	personSalary1["mike"] = 9000
	fmt.Println("personSalary1 map contents:", personSalary1)
	// 2.初始化赋值
	personSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary2["mike"] = 9000
	fmt.Println("personSalary2 map contents:", personSalary2)
	// 根据key获取value, 如果不存在返回0 ??? 这个没问题???
	employee := "nothere"
	fmt.Println("Salary of nothere is", personSalary2[employee])
	// 可以获取到底存不存在, _是value 我们不需要
	_, ok := personSalary2[employee]
	fmt.Println(ok) // false表明不存在
	// for range循环, for range有坑
	// 1方面其不能保证顺序, 2方面其获取的是拷贝值, 无法修改原切片
	for key, value := range personSalary2 {
		fmt.Printf("personSalary2[%s] = %d\n", key, value)
	}
	// 删除
	delete(personSalary2, "steve")
	fmt.Println("steve已经被删除, 现在的切片长度是:", len(personSalary2))
	// ==只能跟nil比, 不能map之间比, 需要自己额外实现
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	if map1 == nil {

	}
	// map转json字符串
	accountPatch := map[string]interface{}{
		"command": "profile_change",
		"effect": map[string]interface{}{
			"mode": "right_now",
		},
		"data": map[string]interface{}{},
	}
	info, _ := json.Marshal(accountPatch)
	fmt.Println(string(info))
}
