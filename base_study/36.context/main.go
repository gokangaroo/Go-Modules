package main

import (
	"context"
	"fmt"
)

func main() {
	// 1.WithCancel例子

	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	//gen := func(ctx context.Context) <-chan int {
	//	dst := make(chan int)
	//	n := 1
	//	go func() {
	//		for {
	//			select {
	//			case <-ctx.Done():
	//				return // returning not to leak the goroutine
	//			case dst <- n:
	//				n++
	//			}
	//		}
	//	}()
	//	return dst
	//}
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // cancel when we are finished consuming integers
	//
	//// range chan
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		break
	//	}
	//}

	// 2.WithDeadLine/WithTimeout例子

	//d := time.Now().Add(50 * time.Millisecond)
	//ctx, cancel := context.WithDeadline(context.Background(), d)
	//ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//
	//// Even though ctx will be expired, it is good practice to call its
	//// cancelation function in any case. Failure to do so may keep the
	//// context and its parent alive longer than necessary.
	//defer cancel()
	//
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err())
	//}

	// 3.WithValue例子,
	//cancel是主动撤销, 而timeout则分为超时和主动撤销(成功请求), 最后value的场景就是传递一个值.但是找不到的值可以在父节点找
	//撤销会传给所有子值, 主要的context分为四种: 根Background, 可撤销的几个,携带值的, 不知道传啥的TODO
	type contextKey string

	f := func(ctx context.Context, k contextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Printf("key:%s's value found:%+v\n", k, v)
			return
		}
		fmt.Printf("key:%s's value not found\n", k)
	}

	k := contextKey("language")
	ctx := context.WithValue(context.Background(), k, "Golang")

	f(ctx, k)
	f(ctx, contextKey("color"))
}
