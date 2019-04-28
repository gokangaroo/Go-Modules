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
func DoAmbiguousQuery(key string, page string, accurateChan chan<- string) {
	// 先拿到json
	jsonStr, _ := GetJson("http://route.showapi.com/1196-1?showapi_appid=91358&showapi_sign=a0eef1cfdd2346c8a4ffe590c50eba4c&keyword=" + key + "&page=" + page + "&rows=10")
	// 将json转化为成语集合, 并放入全局map
	idiomsMap := ParseJson2Idioms(jsonStr)
	for title, idiom := range idiomsMap {
		dbData[title] = idiom
	}
	// 将成语名字写入精确管道
	for title, _ := range idiomsMap {
		accurateChan <- title
	}
	//accurateChan <- "两全十美"
	//accurateChan <- "隔岸观火"
	//accurateChan <- "人猿泰山"
	fmt.Println("DoAmbiguousQuery", ": ", key, "-", page)
}

// 精确查询
func DoAccurateQuery(key string) {
	jsonStr, _ := GetJson("http://route.showapi.com/1196-2?showapi_appid=91358&showapi_sign=a0eef1cfdd2346c8a4ffe590c50eba4c&keyword=" + key + "&page=1&rows=10")
	idiom := ParseJson2Idiom(jsonStr)
	// 完善模糊查询没有的数据
	dbData[idiom.Title] = idiom
	fmt.Println("DoAccurateQuery", ": ", key)
}
