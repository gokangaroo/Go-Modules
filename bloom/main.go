package main

import "fmt"

func main() {
	filter := NewBloomFilter()
	str1 := "bloom one"
	filter.add(str1)
	str2 := "bloom two"
	filter.add(str2)

	fmt.Println(filter.contains(str1))
	fmt.Println(filter.contains("none"))
}
