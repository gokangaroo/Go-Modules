package main

import (
	"fmt"
	"sync"
)

var x = 0

// 1.这个协程就是+1
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// 3.引入mutex锁,并在协程内调用
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	//2.打印1000个协程后的结果, 基本不可能是1000了
	fmt.Println("final value of x", x)
}
