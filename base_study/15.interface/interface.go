package main

import (
	"fmt"
)

//VowelsFinder :定义的接口
type VowelsFinder interface {
	FindVowels() []rune
}

// MyString :...
type MyString string

//FindVowels :实现接口VowelsFinder的MyString连接点
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		// 打印aiueo
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	// 这声明啥意思.
	name := MyString("Sam Anderson")
	// 类似于java多态,编译类型是接口, 实际类型是实现类
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

}
