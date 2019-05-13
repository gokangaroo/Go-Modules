package main

import "fmt"

// Income :项目信息和收入的两个方法
type Income interface {
	calculate() int
	source() string
}

// FixedBilling :项目1, 项目名和项目受注金融
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

// TimeAndMaterial :项目2, 项目名,时薪和工时
type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}

// 结构体实现接口方法, 返回名称和实际收入
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// 打印总收入
func calculateNetIncome(ic []Income) {
	var netincome int
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

// Advertisement :项目3, 广告相关, 新加的项目, 多态特性,无需修改总收入方法
type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}
