package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Math int

type Args struct {
	A int
	B int
}

func (m *Math) Multiply(arg Args, reply *int) error {
	*reply = arg.A * arg.B
	fmt.Println("Multiply 被远程调用")
	return nil
}

type Quotient struct {
	Quo, Rem int
}

func (m *Math) Division(arg Args, quo *Quotient) error {
	if arg.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = arg.A / arg.B
	quo.Rem = arg.A % arg.B
	fmt.Println("Division 被远程调用")
	return nil
}

func main() {
	main := new(Math)

	err := rpc.Register(main)
	if err != nil {
		fmt.Println(err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":20198")
	if err != nil {
		fmt.Println(err)
		return
	}

	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//只需要将连接句柄传入rpc包会自动解析
		rpc.ServeConn(conn)
	}
}
