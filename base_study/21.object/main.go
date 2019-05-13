package main

import (
	"base_study/21.object/employee"
)

func main() {
	// 1.默认构造器的创建方式, 很傻, 怎么使用类似构造器的方式呢?其实就是包装个方法
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }

	// 2. NewT(parameters), 如果只有一个构造器就直接New(parameters),
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
