package main

import (
	"fmt"

	"base_study/13.struct/computer"
)

var _ = computer.Spec{}

func main() {

	//creating structure using field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	//匿名结构体, 跟java匿名内部类一样
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)
	// 通过.进行取出或者改值
	fmt.Println("Employee 3 first_name is", emp3.firstName)

	// 默认是空值和0
	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4)

	//结构体指针
	emp5 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Printf("struct index type is %T\n", emp5) // *main.Employee
	fmt.Println("First Name:", (*emp5).firstName)
	fmt.Println("Age:", emp5.age) // 可以直接调用, 自动解引用

	//字段缺省
	var p Person
	p.string = "爱德华"
	p.int = 23
	p.Address = Address{"无锡市", "新吴区"}
	fmt.Println(p)
	fmt.Println(p.city) //嵌套结构体+匿名字段=提升字段, 也就是不需要p.Address.city

	// 如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。
	// 同样，如果结构体里的字段首字母大写，它也能被其他包访问到。

	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//spec.model = "Mac Mini" //小写就无法访问
	fmt.Println("Spec:", spec)

	//结构体可以直接比较, 但是如果含有不能比较的如map字段, 那么结构体也会无法比较
	var p2 Person
	p2.string = "爱德华"
	p2.int = 23
	p2.Address = Address{"无锡市", "新吴区"}
	fmt.Println(p == p2) //如果都相等就是true
}

// Employee :顾客
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// Person :字段缺省,默认名字就是类型本身
type Person struct {
	string
	int
	Address
}

// Address :地址, 嵌套结构体
type Address struct {
	city, state string
}
