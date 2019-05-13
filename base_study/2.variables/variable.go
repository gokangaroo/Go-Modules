package main

import (
	"fmt"
)

func main() {
	a, b := 20, 30
	fmt.Println("a is", a, "b is", b)
	a, b = 40, 50
	fmt.Println("new a is", a, "new b is", b)
	//a, b := 40, 50
	//a = "nav"

	i := 55   //int
	j := 67.8 //float64
	//sum := i + j
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)
}
