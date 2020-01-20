package utils

import (
	"reflect"
	"unicode/utf8"
)

// 数组相关的, 当然golang用的最多的实际还是切片, 数组是固定死长度的值类型, 而切片更加灵活

// ArrayContains 判断一个数组是否有某个元素,并返回下标, -1表示不存在
func ArrayContains(array interface{}, val interface{}) (index int) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
			}
		}
	default:
		index = -1
	}
	return
}

// ArrayCheckShortest 判断一个数组是否存在某个元素,这个元素的长度不足limit==>标点符号也算
// 如果存在就表示有元素长度不满足, 返回true, 没有就返回false==>场景: 检验班级是否有人特别矮
func ArrayCheckShortest(array interface{}, limit int) (charge bool) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		a := reflect.ValueOf(array)
		length := a.Len()
		if length == 0 {
			return
		}
		for i := 0; i < length; i++ {
			s := ToString(a.Index(i).Interface())
			if utf8.RuneCountInString(s) < limit {
				charge = true
				break
			}
		}
	default:
		charge = false
	}
	return
}

// ArrayCheckLongest 与ArrayCheckShortest相反
func ArrayCheckLongest(array interface{}, limit int) (charge bool) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		a := reflect.ValueOf(array)
		length := a.Len()
		if length == 0 {
			return
		}
		for i := 0; i < length; i++ {
			s := ToString(a.Index(i).Interface())
			if utf8.RuneCountInString(s) > limit {
				charge = true
				break
			}
		}
	default:
		charge = false
	}
	return
}
