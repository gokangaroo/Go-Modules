package main

import "fmt"

// Describer 接口
type Describer interface {
	Describe()
}

// Person :值接收器实现接口
type Person struct {
	name string
	age  int
}

// Describe 使用值接受者Person实现
func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

// Address :指针接收器实现接口
type Address struct {
	state   string
	country string
}

// Describe 使用指针接受者*Address实现
func (a *Address) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	// 1.使用值接受者声明的方法，既可以用值来调用，也能用指针调用
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	// 2.对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
	// d2 = a这句类型不一并不能取到地址会报错, 但a.Describe()却是可以的
	var d2 Describer
	a := Address{"Washington", "USA"}
	//d2 = a
	d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describe()
	a.Describe()

}
