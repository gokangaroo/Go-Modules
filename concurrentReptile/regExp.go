package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// 正则匹配规则
var (
	//rePhone = `1[3456789]\d{9}`
	// 号码存入的同时再存入三个分组
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})`
	//reEmail = `[1-9]\d{4,}@qq\.com`
	//reEmail = `\w+@\w+\.com`
	// 不仅仅局限于qq邮箱和com结尾, \w单词字符, +多个
	reEmail = `\w+@\w+\.[a-z]{2,5}(\.[a-z]{2,5})?`
	// 超链接, hao123就是爬a标签的href, [\s\S]+任意字符多个, ?非贪婪
	reLink = `<a[\s\S]+?href="(http[\s\S]+?)"`
)

// 文档:https://studygolang.com/static/pkgdoc/pkg/regexp.htm
func main() {
	//spiderPhone()
	//spiderEmail()
	spiderLink()
}

func spiderLink() {
	// 抓取超链接
	html := getHtml("http://www.hao123.com")
	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(html, -1)
	for _, x := range results {
		fmt.Println(x[1])
	}
}

func spiderEmail() {
	// 邮箱抓取
	getEmail := getHtml("http://tieba.baidu.com/p/2544042204")
	getEmail += "xxx@saa.com.cn"
	re2 := regexp.MustCompile(reEmail)
	// 返回的是[][]string, -1表示全摘
	allEmail := re2.FindAllStringSubmatch(getEmail, -1)
	// fmt.Println(allString)
	for _, x := range allEmail {
		//fmt.Println(x)
		fmt.Println(x[0])
	}
}

func spiderPhone() {
	// 手机号抓取
	getPhone := getHtml("http://www.taohaoma.com/gujia/")
	// 2.正则匹配
	re1 := regexp.MustCompile(rePhone)
	// 返回的是[][]string, -1表示全摘, 0-n表示抓取n个
	allPhone := re1.FindAllStringSubmatch(getPhone, -1)
	// fmt.Println(allString)
	for _, x := range allPhone {
		fmt.Println(x)
		//fmt.Println(x[0])
	}
}

// 1.获取网页源代码
func getHtml(url string) string {
	resp, err := http.Get(url)
	HandleError(err, `http.Get(url)`)
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, `ioutil.ReadAll(resp)`)
	str := string(bytes)
	// fmt.Println(str)
	return str
}

// 处理异常
func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}
