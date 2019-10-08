package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"runtime"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	var host string
	flag.StringVar(&host, "host", "127.0.0.1", "目标主机")
	flag.Parse()

	client, err := rpc.Dial("tcp", host+":20198")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Println(err.Error(), file, line)
		return
	}

	arg := &Args{5, 5}
	var reply int
	err = client.Call("Math.Multiply", arg, &reply)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Println(err.Error(), file, line)
		return
	}
	fmt.Printf("Math.Multiply: %dx%d=%d\n", arg.A, arg.B, reply)

	var que Quotient
	err = client.Call("Math.Division", arg, &que)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Println(err.Error(), file, line)
		return
	}
	fmt.Printf("Math.Division %d/%d=%d,%d%%%d=%d\n", arg.A, arg.B, que.Quo, arg.A, arg.B, que.Rem)
}
