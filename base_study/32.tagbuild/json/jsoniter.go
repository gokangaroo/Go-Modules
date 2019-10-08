// +build third

package json

import (
	"fmt"
	"github.com/json-iterator/go"
)

var (
	// 这个json用法够奇葩, 一点都不美观
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	fmt.Println("Use [jsoniter] package")
	return json.MarshalIndent(v, prefix, indent)
}
