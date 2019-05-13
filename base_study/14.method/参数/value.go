package main

import "fmt"

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	// 当一个函数有一个值参数，它只能接受一个值参数。
	// 当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p.area() //通过指针调用值接收器

	perimeter(p)
	p.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	r.perimeter() //使用值来调用指针接收器
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}
