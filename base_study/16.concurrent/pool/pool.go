package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

// 需要关闭编译器本身优化: go build -gcflags="-l -N" -o pool
func main() {
	a := time.Now().Unix()
	// 不使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	// 使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}

//without pool  29 s
//with    pool  19 s
//个人尝试只快了三分之一..但是由于Get的锁机制,临时池更多用于高并发场景(13更进行了优化)
