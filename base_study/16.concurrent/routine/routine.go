package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//time.Sleep(1 * time.Second)

	//var count uint32
	// fn 延迟到trigger执行==>按照自增顺序执行, 使用count
	//trigger := func(i uint32, fn func()) {
	//	for {
	//		if n := atomic.LoadUint32(&count); n == i {
	//			fn()
	//			atomic.AddUint32(&count, 1)
	//			break
	//		}
	//		time.Sleep(time.Nanosecond)
	//	}
	//}
	//for i := uint32(0); i < 10; i++ {
	//	go func(i uint32) {
	//		fn := func() {
	//			fmt.Println(i)
	//		}
	//		trigger(i, fn)
	//	}(i)
	//}
	//trigger(10, func() {
	//	fmt.Println("所有协程已执行结束")
	//})

	// 上面方法比较绕, 其实现的是决定协程的执行顺序从而实现顺序打印(本质也是阻塞)
	// 启动十个一模一样的协程, 谁取到数据就打印即可(共享的数据++)
	// 两种本质上都是并发控制, 前者是count, 后者是chan
	var count uint32
	var channel = make(chan uint32, 1)
	for i := 0; i < 10; i++ {
		go func(ch chan uint32) {
			fmt.Println(<-ch)
			count++
			ch <- count
		}(channel)
	}
	channel <- count
	time.Sleep(time.Second)
}
