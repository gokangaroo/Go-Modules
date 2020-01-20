package utils

import (
	"encoding/binary"
	"reflect"
	"strconv"
	"strings"
)

// ParseInt64 通用类型转int64, 实现方法与ToString有一定区别
func ParseInt64(val interface{}) (d int64) {
	switch v := val.(type) {
	case string:
		if strings.Count(v, ".") == 1 {
			f, _ := strconv.ParseFloat(v, 64)
			d = ParseInt64(f)
		} else {
			d, _ = strconv.ParseInt(v, 10, 64)
		}
	case []byte:
		d = int64(binary.BigEndian.Uint64(v))
	case int, int8, int16, int32, int64:
		d = reflect.ValueOf(val).Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(reflect.ValueOf(val).Uint())
	case bool:
		if v {
			d = 1
		} else {
			d = 0
		}
	case float32, float64:
		d = int64(reflect.ValueOf(val).Float())
	default:
		d = 0
	}
	return
}
