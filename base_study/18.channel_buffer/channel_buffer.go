package main

import (
	"fmt"
)

func main() {
	//1.缓冲信道, 容量为2, 在没有满/空的时候并不会产生阻塞
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	//ch <- "死锁了"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println(<-ch)
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println(<-ch)
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	// 缓冲信道有点类似队列, FIFO, 先进先出
}
