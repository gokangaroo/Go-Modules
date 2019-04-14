package demoApp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// 成语结构体
type Idiom struct {
	Title      string
	Spell      string
	Content    string
	Sample     string
	Derivation string
}

// 定义全局变量,默认是相对路径
var idiomsMap map[string]Idiom
//var dstFilePath = "f:/go_space/src/idiomApp/成语大全.json"
var dstFilePath = "成语大全.json"

// 获得数据
func init() {
	idiomsMap = make(map[string]Idiom)
	// 3,数据的本地读取和持久化
	fileInfo, e := os.Stat(dstFilePath)
	if fileInfo != nil && e == nil {
		e := ReadIdiomsFromFile(dstFilePath)
		if e != nil {
			defer WriteIdioms2File(dstFilePath)
		} else {
			return
		}
	} else {
		defer WriteIdioms2File(dstFilePath)
	}
	// 1,获取数据
	jsonStr, _ := GetJson("http://route.showapi.com/1196-1?showapi_appid=91358&showapi_sign=a0eef1cfdd2346c8a4ffe590c50eba4c&keyword=肉&page=1&rows=10")
	//fmt.Println(jsonStr)
	// 2,解析json为成语map
	ParseJson2Idioms(jsonStr)
	//fmt.Println(idiomsMap)
	// 4,查询已有title(除了成语)的详细解释(api)并解析替换
	AddDetails(idiomsMap)
}

func main() {
	// 4,获取命令行参数,保证好的包GetCmdlineArgs
	// idiom.exe -cmd ambiguous -keyword 肉
	// idiom.exe -cmd accurate -keyword 不知肉味
	// go build -o idiom.exe ./
	cmdInfo := [3]interface{}{"cmd", "未知命令", "ambiguous=模糊查询, accurate=精确查询"}
	keywordInfo := [3]interface{}{"keyword", "未知关键字", "查询关键字"}
	retValuesMap := GetCmdlineArgs(cmdInfo, keywordInfo)
	//fmt.Println(retValuesMap)

	//5. 执行查询
	cmd := retValuesMap["cmd"].(string)
	keyword := retValuesMap["keyword"].(string)
	if cmd == "ambiguous" {
		// 模糊查询
		for title, idiom := range idiomsMap {
			if strings.Contains(title, keyword) {
				printIdiom(idiom)
			}
		}
	} else if cmd == "accurate" {
		idiom := idiomsMap[keyword]
		printIdiom(idiom)
	} else {
		fmt.Println("非法命令: ", cmd)
	}
}

// 打印成语
func printIdiom(idiom Idiom) {
	if idiom.Title != "" {
		fmt.Println("成语名: ", idiom.Title)
		fmt.Println("读音: ", idiom.Spell)
		fmt.Println("详细解释: ", idiom.Content)
		fmt.Println("使用实例: ", idiom.Sample)
		fmt.Println("出处: ", idiom.Derivation)
		fmt.Println()
	}
}

//5. 根据title添加相关详情
func AddDetails(idiomsMap map[string]Idiom) {
	for title, _ := range idiomsMap {
		// api请求1/1s
		time.Sleep(1 * time.Second)
		jsonStr, err := GetJson("http://route.showapi.com/1196-2?showapi_appid=91358&showapi_sign=a0eef1cfdd2346c8a4ffe590c50eba4c&keyword=" + title + "&page=1&rows=10")
		if err != nil {
			fmt.Println("查询详情失败: ", title)
			continue
		}
		//jsonStr转为idiom对象
		var idiom = Idiom{Title: title}
		var tmp map[string]interface{}
		json.Unmarshal([]byte(jsonStr), &tmp)
		tmp = tmp["showapi_res_body"].(map[string]interface{})["data"].(map[string]interface{})
		idiom.Spell = tmp["spell"].(string)
		idiom.Sample = tmp["samples"].(string)
		idiom.Derivation = tmp["derivation"].(string)
		idiom.Content = tmp["content"].(string)
		//替换之前的Idiom
		idiomsMap[title] = idiom
		fmt.Printf("%s的详情已更新\n", title)
	}
}

// 4.从已有json读进来
func ReadIdiomsFromFile(dstFile string) error {
	//文件已存在, 读入json数据
	//fmt.Println("数据已就绪!")
	file, _ := os.OpenFile(dstFile, os.O_RDONLY, 0666)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&idiomsMap)
	if err != nil {
		fmt.Println("加载数据失败: ", err)
		return err
	}
	fmt.Println("数据加载成功!")
	return nil
}

// 3.数据持久化, 存入json文件(这个文件需要先去创建)
func WriteIdioms2File(dstFile string) {
	//os.Create(dstFile)
	//没有就创建, 有就覆盖
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

// 2.解析json为成语map
func ParseJson2Idioms(jsonStr string) {
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
}

// 1.获取HTTP报文
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
