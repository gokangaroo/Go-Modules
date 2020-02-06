package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 账户初始值为0元
	var balance int32
	balance = int32(0)
	done := make(chan bool)
	// 执行10000次转账，每次转入1元
	count := 10000

	var lock sync.Mutex
	s := time.Now()
	for i := 0; i < count; i++ {
		// 这里模拟异步并发转账
		go transfer(&balance, 1, done, &lock) //30ms
		//go transferCas(&balance, 1, done)//29ms
		//go transferFaa(&balance, 1, done)//26ms
	}
	// 等待所有转账都完成
	for i := 0; i < count; i++ {
		<-done
	}
	// 打印账户余额
	fmt.Printf("balance = %d ,cost time:%+v \n", balance, time.Since(s))
}

// 转账服务,锁实现
func transfer(balance *int32, amount int, done chan bool, lock *sync.Mutex) {
	lock.Lock()
	*balance = *balance + int32(amount)
	lock.Unlock()
	done <- true
}

// cas原语=>更加通用
func transferCas(balance *int32, amount int, done chan bool) {
	for {
		old := atomic.LoadInt32(balance)
		new := old + int32(amount)
		if atomic.CompareAndSwapInt32(balance, old, new) {
			break
		}
	}
	done <- true
}

// faa原语
func transferFaa(balance *int32, amount int, done chan bool) {
	atomic.AddInt32(balance, int32(amount))
	done <- true
}
