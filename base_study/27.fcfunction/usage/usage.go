package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func main() {
	// 1.创建学生切片
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	// 其中传入的函数参数是一个匿名函数, 定义的条件是grade == "B"
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}

// 这个函数计算了某一学生是否满足筛选条件。
func filter(s []student, f func(student) bool) []student {
	var r []student
	// 表扫描, 扫到放到结果集
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}
