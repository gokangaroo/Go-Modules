package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello World"
	printBytes(name)
	printChars(name)
	name = "你好"
	printBytes(name)
	printChars(name) // 这里发生了了编码错误,一个汉字占用3个字节
	//rune 能帮我们解决这个难题, 转换后遍历
	printTrueChars(name)
	//for range遍历, 可以直接使用rune接受
	printCharsAndBytes(name)
	// 字节切片构造字符串, 16进制, 10进制都可以
	byteSlice := []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd}
	str := string(byteSlice)
	fmt.Println(str)
	// rune切片构造字符串, 但是中文会乱码???
	// {0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} 这个就不会乱码
	runeSlice := []rune{0x00e4, 0x00bd, 0x00a0, 0x00e5, 0x00a5, 0x00bd}
	strrune := string(runeSlice)
	fmt.Println(strrune)
	// 字符串长度, 会按照rune的长度来计算, 所以中文, 本来是3个字节, 但是在这里因为占用2个rune所以是2
	fmt.Println(utf8.RuneCountInString(str))
	//改成"我好"
	fmt.Println(mutate([]rune(str)))
}

// 循环输出字符串字符 16 进制编码
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) //%x 格式限定符用于指定 16 进制编码。
	}
	fmt.Println()
}

// 循环输出每个字符, 但是有中文会乱码
func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
}

// 循环输出每个字符, 使用rune
func printTrueChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}

// for range循环
func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

//改成"我好"
func mutate(s []rune) string {
	s[0] = '我'
	return string(s)
}
