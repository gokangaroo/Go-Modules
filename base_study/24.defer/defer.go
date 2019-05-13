package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	nums := []int{78, 109, 2, 563, 300}
	// defer接函数
	largest(nums)
	// defer接方法尝试, 也是可以的
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	// 会在函数结束前, 调用这个方法
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}
