package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 将json解析成map
// one:模糊查询
func ParseJson2Idioms(jsonStr string) (idiomsMap map[string]Idiom) {
	idiomsMap = make(map[string]Idiom)
	// 将json改为go的数据, 随便用个interface接受一下.
	tempData := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &tempData)
	//fmt.Println(tempData)
	// 挖出tempData有用的数据放到结构体中, map里面也是map, 需要类型断言
	tempSlice := tempData["showapi_res_body"].(map[string]interface{})["data"].([]interface{})
	//fmt.Printf("type=%T,value=%v", tempSlice, tempSlice)
	// 接口切片导入结构体
	for _, v := range tempSlice {
		// 类型断言, key value都是string, value拿出来再断言一次string
		title := v.(map[string]interface{})["title"].(string)
		// 都放到map里面, title做key, idiom本身做值
		idiom := Idiom{Title: title}
		idiomsMap[title] = idiom
	}
	//fmt.Println(idiomsMap)
	return
}

// two: 精确查询
func ParseJson2Idiom(jsonStr string) (idiom Idiom) {
	/*TODO*/
	return Idiom{}
}

// json的文件持久化与读取
func WriteIdioms2File(idiomsMap map[string]Idiom, dstFile string) {
	//os.Create(dstFile)
	//没有就创建, 有就覆盖
	/*TODO*/
	file, _ := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(idiomsMap)
	if err != nil {
		fmt.Println("数据写入Json文件失败: ", err)
		return
	}
	fmt.Println("数据写入Json文件成功!")
}

// 读入json为map
func ReadIdiomsFromFile(dstFile string) (idiomsMap map[string]Idiom, err error) {
	idiomsMap = make(map[string]Idiom)
	//文件已存在, 读入json数据
	//fmt.Println("数据已就绪!")
	file, _ := os.OpenFile(dstFile, os.O_RDONLY, 0666)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&idiomsMap)
	if err != nil {
		fmt.Println("加载数据失败: ", err)
	} else {
		fmt.Println("数据加载成功!")
	}
	return
}
