package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var ff, e float64
	e = 100.00
	ff = -0.210615789

	// 先计算
	qq := ff * e
	fmt.Println(qq) // 输出 -21.0615789
	// 再截取
	qq = FloatRound(qq, 4)
	fmt.Println(qq) // 输出 -21.0616

	i1 := time.Now().UnixNano()
	time.Sleep(1 * time.Second)
	i2 := time.Now().UnixNano()
	f := float64(i2-i1) / 1e6
	fmt.Println(FloatRound(f, 2))
	f = float64(111)
	fmt.Println(f)
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
