package main

import "fmt"

func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
	fmt.Println("length of a is", len(a))
	fmt.Printf("====================\n")

	sum := "countries:"
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %s\n", i, v)
		sum += v
	}
	fmt.Println(sum)
	fmt.Printf("====================\n")

	c := [3][2]string{ //3行两列
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	for _, v1 := range c { //嵌套循环
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("====================\n")

	d := c[0:1]            //切片多维数组多行怎么切?好像只能一行一行切
	for _, v1 := range d { //嵌套循环
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("====================\n")

	f := make([][]string, 3) //三行,容量3的多维切片,容量不指定就跟长度相等
	for i := 0; i < len(f); i++ {
		f[i] = make([]string, 3, 5) //每行长度3,容量5
	}
	//多维切片初始化,切片赋值等同于前面一维切片赋值,只能一个一个赋值了
	for x1, v1 := range f { //嵌套循环
		fmt.Printf("这是第%d行: ", x1)
		for x2, v2 := range v1 {
			//然后可以租个赋值
			fmt.Printf("这是第%d列 ", x2)
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("当前切片长度为%d\n", len(f))
	fmt.Printf("当前切片容量为%d\n", cap(f))
	fmt.Printf("====================\n")
	e := make([]string, 1)
	f = append(f, e)
	for x1, v1 := range f { //嵌套循环
		fmt.Printf("这是第%d行: ", x1)
		for x2, v2 := range v1 {
			fmt.Printf("这是第%d列 ", x2)
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("当前切片长度为%d\n", len(f))
	fmt.Printf("当前切片容量为%d\n", cap(f))
	fmt.Printf("====================\n")
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Printf("food:%s\n", food)
	fmt.Printf("====================\n")

	countries := [5]string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Println(countriesCpy)           //这样countries就会被回收
	fmt.Printf("====================\n")
}
