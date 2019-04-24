package main

import (
	"flag"
	"fmt"
)

/*
获取用户从命令行输入的参数值，以map[string]interface{}形式返回
argInfos：其中的每个元素为一个三长度的数组，分别放置一个命令的[名称 默认值 用法说明]
retValuesMap：retValuesMap["cmd"]即为用户在命令行中输入的cmd参数值，其它的以此类推
*/
func GetCmdlineArgs(argInfos ...[3]interface{}) (retValuesMap map[string]interface{}) {

	fmt.Printf("type=%T,value=%v\n", argInfos, argInfos)

	//初始化返回结果
	retValuesMap = map[string]interface{}{}

	//预定义【用户可能输入的各种类型的指针】
	var strValuePtr *string
	var intValuePtr *int

	//预定义【用户可能输入的各种类型的指针】的容器
	//用户可能输入好几个string型的参数值，存放在好几个string型的指针中，将这些同种类型的指针放在同种类型的map中
	//例如：flag.Parse()了以后，可以根据【strValuePtrsMap["cmd"]】拿到【存放"cmd"值的指针】
	var strValuePtrsMap = map[string]*string{}
	var intValuePtrsMap = map[string]*int{}

	/*	var floatValuePtr *float32
		var floatValuePtrsMap []*float32
		var boolValuePtr *bool
		var boolValuePtrsMap []*bool*/

	//遍历用户需要接受的所有命令定义
	for _, argArray := range argInfos {

		/*
		先把每个命令的名称和用法拿出来,
		这俩货都是string类型的，所有都可以通过argArray[i].(string)轻松愉快地获得其字符串
		一个叫“cmd”，一个叫“你想干嘛”
		"cmd"一会会用作map的key
		*/
		//[3]interface{}
		//["cmd" "未知类型" "你想干嘛"]
		//["gid"     0     "要查询的商品ID"]
		//上面的破玩意类型[string 可能是任意类型 string]
		nameValue := argArray[0].(string)  //拿到第一个元素的string值,是命令的name
		usageValue := argArray[2].(string) //拿到最后一个元素的string值，是命令的usage

		//判断argArray[1]的具体类型
		switch argArray[1].(type) {
		case string:
			//得到【存放cmd的指针】，cmd的值将在flag.Parse()以后才会有
			//cmdValuePtr = flag.String("cmd", argArray[1].(string), "你想干嘛")
			strValuePtr = flag.String(nameValue, argArray[1].(string), usageValue)

			//将这个破指针以"cmd"为键，存在【专门放置string型指针的map，即strValuePtrsMap】中
			strValuePtrsMap[nameValue] = strValuePtr

		case int:
			//得到【存放gid的指针】，gid的值将在flag.Parse()以后才会有
			//gidValuePtr = flag.String("gid", argArray[1].(int), "商品ID")
			intValuePtr = flag.Int(nameValue, argArray[1].(int), usageValue)

			//将这个破指针以"gid"为键，存在【专门放置int型指针的map，即intValuePtrsMap】中
			intValuePtrsMap[nameValue] = intValuePtr
		}

	}

	/*
	程序运行到这里，所有不同类型的【存值指针】都放在对相应类型的map中了
	flag.Parse()了以后，可以从map中以参数名字获取出【存值指针】，进而获得【用户输入的值】
	*/

	//用户输入完了，解析，【用户输入的值】全都放在对应的【存值指针】中
	flag.Parse()

	/*
	遍历各种可能类型的【存值指针的map】
	*/
	if len(strValuePtrsMap) > 0 {
		//从【cmd存值指针的map】中拿取cmd的值，还以cmd为键存入结果map中
		for k, vPtr := range strValuePtrsMap {
			retValuesMap[k] = *vPtr
		}
	}
	if len(intValuePtrsMap) > 0 {
		//从【gid存值指针的map】中拿取gid的值，还以gid为键存入结果map中
		for k, vPtr := range intValuePtrsMap {
			retValuesMap[k] = *vPtr
		}
	}

	//返回结果map
	return
}
