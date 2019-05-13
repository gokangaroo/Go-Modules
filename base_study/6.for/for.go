package main

import (
	"fmt"
)

func main() {
	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}
}
