package main

import (
	"fmt"
)

// SalaryCalculator :薪资计算接口
type SalaryCalculator interface {
	CalculateSalary() int
}

// Permanent :长期员工,薪资=basicpay+pf
type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

// Contract :合同员工,薪资=basicpay
type Contract struct {
	empID    int
	basicpay int
}

//CalculateSalary :长期员工,薪资= basicpay + pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//CalculateSalary :合同员工,薪资= basicpay
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// 传入一个接口数组切片.获取所有员工的薪资综合
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	// 5000+6000+3000+20+30=14050
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}
