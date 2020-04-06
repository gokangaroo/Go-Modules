package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			add("https://github.com/huija")
		}
	}()

	//http://127.0.0.1:6060/debug/pprof/ #在页面生成profile文件
	//go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60 #命令行也可以
	//go tool pprof -http=:8080 cpu.prof #可视化界面
	http.ListenAndServe("0.0.0.0:6060", nil)
}

var res []string

func add(str string) string {
	data := []byte(str)
	sData := string(data)
	res = append(res, sData)

	return sData
}
