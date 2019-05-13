//Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	// 相当于启动一个子线程,主线程结束他的输出也不会显示出来
	// 哦, go里面叫协程, goroutine,其实就是应用层的线程.
	go hello()
	// 停留1s等待hello协程运行完
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// 再来几个例子
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
