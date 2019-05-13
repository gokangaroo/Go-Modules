package main

import (
	"base_study/22.inherit/inherit"
)

func main() {
	// 方法全是大写开头, 类似public, 结构体定义都是小写开头, 用构造器传出来
	author1 := inherit.Newauthor("Naveen", "Ramanathan", "Golang Enthusiast")
	post1 := inherit.Newpost("Inheritance in Go", "Go supports composition instead of inheritance", author1)
	//post1.Details()
	post2 := inherit.Newpost("Struct instead of Classes in Go", "Go does not support classes but methods can be added to structs", author1)
	post3 := inherit.Newpost("Concurrency", "Go is a concurrent language and not a parallel one", author1)
	website := inherit.Newwebsite(inherit.Newposts(post1, post2, post3))
	website.Contents()

}
