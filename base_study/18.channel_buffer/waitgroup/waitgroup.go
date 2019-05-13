package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	// 协程执行完就Done, 相当于-1
	wg.Done()
}

func main() {
	no := 3
	//WaitGroup 用于等待一批 Go 协程执行结束
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		// 每启动一个协程我们就+1
		wg.Add(1)
		go process(i, &wg)
	}
	// 然后通wait,等待所有协程执行完
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
