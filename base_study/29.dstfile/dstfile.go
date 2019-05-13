package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// ioutil.ReadFile
	// 1.程序内写死的相对路径或者绝对路径
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	// 2.flag读取命令行参数, 输入路径
	// go build -o dstfile.exe ./
	// dstfile -fpath=f:/go_space/src/base_study/29.dstfile/test.txt
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	// 参数填充, 上面的"test.txt"是默认值, 最后一个一般是usage
	flag.Parse()
	fmt.Println(*fptr)
	data, err = ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	// 3.分块读取大文件,使用 bufio 包来完成
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			//直到读到EOF结尾
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
	// bufio包还可以创建Scanner对象逐行读取

}
