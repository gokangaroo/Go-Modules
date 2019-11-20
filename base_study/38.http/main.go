package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello guy!") // 这个写入到 w 的是输出到客户端的
}

func main() {
	//type ServeMux struct {
	//	mu    sync.RWMutex
	//	m     map[string]muxEntry
	//	es    []muxEntry // slice of entries sorted from longest to shortest.
	//	hosts bool       // whether any patterns contain hostnames
	//}
	http.HandleFunc("/", sayhelloName) // 设置访问的路由, ServeMux, golang默认的路由器, map进行匹配
	//		c := srv.newConn(rw)
	//		c.setState(c.rwc, StateNew) // before Serve can return
	//		go c.serve(ctx)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口,每来一个请求就建立一个连接,然后抛给具体的路由进行处理
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
