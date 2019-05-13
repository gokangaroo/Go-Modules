package main

import "fmt"

// Describer :接口
type Describer interface {
	Describe()
}

// Person :连接点
type Person struct {
	name string
	age  int
}

// Describe :实现方法签名
func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer: //直接跟接口比
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType(p)
}
