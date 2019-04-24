package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 从url获取到json数据
func GetJson(url string) (jsonStr string, err error) {
	// 1.获得数据
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http请求失败: ", err)
		return
	}
	// 0.延时关闭HTTP连接
	defer resp.Body.Close()
	// 2.resp.Body实现了Reader接口, 可以进行字节读取.
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("报文读取失败: ", err)
		return
	}
	// 3.Http文本(Json, html...)转为字符串
	jsonStr = string(bytes)
	return
}

// 模糊查询
func DoAmbiguousQuery(key string, page string, accurateChan chan<- string) (jsonStr string) {
	/*TODO*/
	accurateChan <- "两全十美"
	accurateChan <- "隔岸观火"
	accurateChan <- "人猿泰山"
	fmt.Println("DoAmbiguousQuery", ": ", key, "-", page)
	return ""
}

// 精确查询
func DoAccurateQuery(key string) (jsonStr string) {
	/*TODO*/
	fmt.Println("DoAccurateQuery", ": ", key)
	return ""
}
