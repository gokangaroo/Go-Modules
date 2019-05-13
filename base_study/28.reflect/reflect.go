package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}
type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	//创建一个结构体, 现在我们要将这个结构体insert 数据库!
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQueryDemo(o))
	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)

}

// 创建sql的方法
func createQueryDemo(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

// 现在我们要使得他们变得通用, 那么就得用Object类型, 哦不, Go是interface{}
func createQuery(q interface{}) {
	// 但是结构体都不知道有几个参数,有什么类型, 怎么来实现呢???
	// 得到结构体对象
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name() // t是结构体名称, 这里对应db的表名
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)             // v是结构体对象本身
		for i := 0; i < v.NumField(); i++ { //遍历属性
			switch v.Field(i).Kind() { //查看属性的Kind
			case reflect.Int: //可以点进去看reflect支持多少类型
				if i == 0 { //第一个就不需要逗号,直接拼接在原query后面
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		// 给query收一个尾巴, 括号补上
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
	// reflect 包会帮助识别 interface{} 变量的底层具体类型(全限定名)和具体值(对象): reflect.TypeOf() 和 reflect.ValueOf()
	// 还有Kind属性,可以返回struct这种, 大而化之的类型--> calss, method, 还是field
	// NumField() 方法返回结构体中字段的数量，而 Field(i int) 方法返回字段 i 的 reflect.Value
}
