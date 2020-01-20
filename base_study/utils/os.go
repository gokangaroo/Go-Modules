package utils

import "os"

// CheckFileIfExist 判断文件是否存在
func CheckFileIfExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
