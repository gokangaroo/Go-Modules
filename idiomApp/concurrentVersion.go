package main

import (
	"fmt"
	"os"
	"time"
)

/**
编译命令:
启动命令: idiom.exe -cmd start -poem 两岸猿声啼不住
将诗句的每个字丢入[模糊查询channel]
ext建立[精确查询channel和退出channel]
时钟-->每秒随机选择一条管道取值, 进行相应操作.
	模糊管道: 起协程进行模糊查询, 数据保存入内存
	精确查询管道: 起协程进行精确查询, 数据保存入内存
	退出管道: 所有数据持久化
 */

//0. cmdUtil
//1. 三条chan
//2. 网络解析持久化: models+netUtil+jsonUtil
var (
	ambiguousChan = make(chan string, 20)
	accurateChan  = make(chan string, 20)
	quitChan      = make(chan string, 0)

	// 全局map
	dbData = make(map[string]Idiom)
	// 全局json文件
	jsonFile = "db.json"
)

// 直接运行会undefined, 需要右键包build运行
func main() {
	// 读入命令行参数
	// idiom -cmd start -poem 两岸猿声啼不住
	cmdInfo := [3]interface{}{"cmd", "未知命令", "开始游戏!"}
	poemInfo := [3]interface{}{"poem", "未知命令", "一行诗句!"}
	retValuesMap := GetCmdlineArgs(cmdInfo, poemInfo)
	cmd := retValuesMap["cmd"].(string)
	poem := retValuesMap["poem"].(string)
	fmt.Println(cmd, poem)
	//poem := "两岸猿声啼不住"
	// 使用rune[]或者fmt的unicode字符处理
	// fmt.Printf("%c,%d", 22823, '大')
	// 将诗句打碎丢入模糊管道
	for _, key := range poem {
		ambiguousChan <- fmt.Sprintf("%c", key)
	}
	// 每秒chan三选一, quitChan读到Over就直接结束了.
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			select {
			case key := <-ambiguousChan:
				go DoAmbiguousQuery(key, "1", accurateChan)
			case key := <-accurateChan:
				go DoAccurateQuery(key)
			case <-quitChan:
				WriteIdioms2File(dbData, jsonFile)
				os.Exit(0)
			}
		}
	}()
	timer := time.NewTimer(20 * time.Second)
	// 倒计时20s结束主程序
	<-timer.C
	quitChan <- "OVER"

}
