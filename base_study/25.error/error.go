package main

import (
	"fmt"
	"net"
)

func main() {
	// 1.文件打开错误
	// f, err := os.Open("/test.txt")
	// // 1.只使用error接口太直接
	// // 报错: open /test.txt: The system cannot find the file specified.
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	return
	// // }
	// // 2.使用PathError结构体,可以拼接打印需要的信息
	// // 报错: File at path /test.txt failed to open
	// if err, ok := err.(*os.PathError); ok {
	// 	fmt.Println(err.Op, err.Path, err.Err)
	// 	return
	// }
	// fmt.Println(f.Name(), "opened successfully")

	// 2.域名访问错误,DNSError,可以查看时超时还是暂时的错误
	addr, err := net.LookupHost("golangbot123456.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}
