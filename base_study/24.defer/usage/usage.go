package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	// 使用WaitGroup来等待子协程执行完毕
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// 计算矩形面积
func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		//wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		//wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	//这里针对每一个返回都要使用wg.Done()
	//但是可以在方法开头写defer, 来进行代码简化
	//wg.Done()
}
