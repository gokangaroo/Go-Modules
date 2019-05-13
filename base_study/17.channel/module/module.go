package main

import (
	"fmt"
)

// 拆解数字, 并放入信道
func digits(number int, dchnl chan<- int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

// 调用拆解数字, 计算平方和后放入信道
func calcSquares(number int, squareop chan<- int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

// 调用拆解数字, 计算立方和后放入信道
func calcCubes(number int, cubeop chan<- int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 123
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("平方和:", squares, "\n立方和:", cubes)
}
