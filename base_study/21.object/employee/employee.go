package employee

import (
	"fmt"
)

// Employee :顾客, 改为小写, 禁止其他包直接引用, 必须调用new方法
// type Employee struct {
type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// New :暴露出来的构造器
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

// LeavesRemaining :打印详情
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
