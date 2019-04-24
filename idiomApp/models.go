package main

import "fmt"

// 成语的数据模型
// 成语结构体
type Idiom struct {
	Title      string
	Spell      string
	Content    string
	Sample     string
	Derivation string
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