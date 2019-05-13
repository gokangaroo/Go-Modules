package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {
	f, _ := os.Create("test.txt")
	defer f.Close()
	// 1.写入字符串
	l, _ := f.WriteString("Hello World, String\n")
	fmt.Println(l, "bytes written successfully")
	// 2.写入字节, ASCII吗表:https://baike.baidu.com/item/ASCII/309296?fr=aladdin
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 10}
	n2, _ := f.Write(d2)
	fmt.Println(n2, "bytes written successfully")
	// 3.写入字符串切片, Fprintln-->行添加
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	for _, v := range d {
		fmt.Fprintln(f, v)
	}
	fmt.Println(d, "bytes written successfully")
	// 4.OpenFile,第三个参数FileMode
	f2, _ := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	newLine := "File handling is easy."
	defer f2.Close()
	fmt.Fprintln(f2, newLine)
	// unix的权限可以看作8进制, 而我们是10进制传入,需要前面加0, 或者使用标准库中的strconv.ParseInt
	um, _ := strconv.ParseInt(strconv.Itoa(777), 8, 0)
	fmt.Println(os.FileMode(777), 777)
	fmt.Println(os.FileMode(0777), 0777)
	fmt.Println(os.FileMode(um), um)
	// 5.并发写文件
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	// 创建100个produce协程
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	// 消费100个随机数, 写入文件
	go consume(data, done)
	// 最后进行关闭
	go func() {
		wg.Wait()
		close(data)
	}()
	j := <-done
	if j == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}

// 生产随机数的函数
func produce(data chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := rand.Intn(999)
	data <- n
}

// 将随机数写入文件
func consume(data chan int, done chan bool) {
	f, _ := os.Create("concurrent.txt")
	defer f.Close()
	// 循环读取data的值
	for d := range data {
		fmt.Fprintln(f, d)
	}
	done <- true
}
