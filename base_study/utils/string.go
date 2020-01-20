package utils

import (
	"fmt"
	"strings"
)

// ToString 通用类型转字符串
func ToString(val interface{}) (s string) {
	switch v := val.(type) {
	case string:
		s = strings.TrimSpace(v)
	case uint64, int64, int, int32, uint32, int8, uint8:
		s = fmt.Sprintf("%d", v)
	case float64, float32:
		s = strings.Trim(fmt.Sprintf("%f", v), "0")
	case bool:
		if v {
			s = "true"
		} else {
			s = "false"
		}
	case error:
		s = v.Error()
	default:
		// %v will return the `String()` of the type that satisfy the fmt.Stringer interface.
		// https://stackoverflow.com/questions/41159169/why-not-use-v-to-print-int-and-string
		s = fmt.Sprintf("%+v", v)
	}
	return
}
