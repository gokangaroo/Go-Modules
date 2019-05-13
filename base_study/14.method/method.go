package main

import (
	"fmt"
)

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
		address: address{
			"无锡市",
			"新吴区",
		},
	}
	emp1.displaySalary1() // 通过Employee 类型变量调用displaySalary() 方法
	displaySalary2(emp1)  // 函数是这么用
	// 目的
	// 1.基于类型的方法是一种实现和类相似行为的途径。
	// 2.相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
	emp1.changeName("爱德华")
	emp1.changeSalary(10000)
	emp1.displaySalary1()
	// 直接调用匿名字段的方法
	emp1.fullAddress()
	// 难点,

}

// Employee :顾客
type Employee struct {
	name     string
	salary   int
	currency string
	address
}

type address struct {
	city  string
	state string
}

/*
 方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
 接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
*/
/*
 displaySalary() 方法将 Employee 做为接收器类型, 通过接收器调用方法
*/
func (e Employee) displaySalary1() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
	fmt.Println()
}

// 原函数调用
func displaySalary2(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
	fmt.Println()
}

/*
使用值接收器的方法。
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *Employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

// person可以直接调用匿名字段的方法
func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}
