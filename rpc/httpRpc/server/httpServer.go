package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

//- Registry 的方法要求
//- 方法是导出的
//- 方法有两个参数，都是导出类型或内建类型
//- 方法的第二个参数是指针
//- 方法只有一个error接口类型的返回值

type Math int

type Args struct {
	A int
	B int
}

// Multiply 加法
func (m *Math) Multiply(arg Args, reply *int) error {
	*reply = arg.A * arg.B
	fmt.Println("Multiply 被远程调用")
	return nil
}

type Quotient struct {
	Quo, Rem int
}

// Division 除法,取余
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

	math := new(Math)

	err := rpc.Register(math)
	if err != nil {
		fmt.Println(err.Error())
	}

	rpc.HandleHTTP()

	err = http.ListenAndServe(":20198", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
