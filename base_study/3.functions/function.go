package main

import (
	"fmt"
)

// main
func main() {
	price, no := 90, 6 // 定义 price 和 no,默认类型为 int
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice) // 打印到控制台上
	area1, perimeter1 := rectProps1(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area1, perimeter1)
	area2, _ := rectProps2(10.8, 5.6) //空白符
	fmt.Printf("Area %f", area2)
}

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps1(length, width float64) (float64, float64) { //返回矩形的面积和周长
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
