//所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
package main

import (
	"fmt"
	"time"
)

func main() {
	// 0.声明一个只能传int的chan, 默认是nil,也可以进行简短声明
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
	}
	// 1.通过信道进行发送和接受
	done := make(chan bool)
	go hello(done)
	fmt.Println("小明: 我在等手抓饼")
	// 接受信道的值
	rs := <-done
	fmt.Println(rs, "小明: 好")

}

func hello(done chan bool) {
	// 阻塞特性: 当没有数据存入, 主协程的接受会一直阻塞
	fmt.Println("大妈: 你等会啊,4秒钟")
	time.Sleep(4 * time.Second)
	// 存入信道
	done <- true
	fmt.Println("大妈: 我做好了你拿走吧")
}
