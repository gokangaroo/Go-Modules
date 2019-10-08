// +build !third

//必须留一行空格，否则会报重定义的错误
package json

import (
	"encoding/json"
	"fmt"
)

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	fmt.Println("Use [encoding/json] package")
	//MarshalIndent类似Marshal但会使用缩进将输出格式化。
	return json.MarshalIndent(v, prefix, indent)
}
