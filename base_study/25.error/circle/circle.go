package main

import (
	"fmt"
	"math"
)

// 2.自定义异常
type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

// 1.这里都是简单的输入字符串自定义错误, 还有结构体类型的错误, 需要实现error的Error()方法
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("Area calculation failed, radius is less than zero")
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := 20.0
	a, e := circleArea(radius)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("Area of circle %0.2f\n", a)

	// 自定义异常并使用
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

// 3.实现error接口并自定义更多方法
func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

// 矩形计算面积的方法,返回结果和异常
func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}
