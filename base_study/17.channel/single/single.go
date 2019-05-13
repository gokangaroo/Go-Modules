package main

import (
	"fmt"
)

func sendData(sendch chan<- int) {
	for i := 0; i < 10; i++ {
		sendch <- i
	}
	// 关闭信道,告诉接收方再有信息
	close(sendch)
}

func main() {
	cha1 := make(chan int)
	go sendData(cha1)
	for {
		v, ok := <-cha1
		if ok == false {
			// 如果信道关闭, ok就是false
			fmt.Println("Received ", ok)
			break
		}
		fmt.Println("Received ", v, ok)
	}
	// 使用for range会自动退出, 如果信道关闭.
	// for v := range cha1 {
	// 	fmt.Println("Received ", v)
	// }
}
