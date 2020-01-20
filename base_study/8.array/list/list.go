package main

import (
	"base_study/utils"
	"container/list"
	"container/ring"
	"fmt"
	"log"
	"reflect"
)

func main() {
	// list包实现了双向链表==>实际内部就是个环
	// 有个root点,一直是nil, 可以看做世界的开始, 世界的结束
	// 右边为next, 左边为prev, 右边终点连接左边起点
	// 插入front就是插在root右手, back就是插在root左手
	// 正向遍历就是顺时针往右遍历, 直到遍历到root=nil
	// 反向遍历则相反
	var listDemo list.List
	// 插入操作会懒初始化==>root.next=nil也就是root.next=root的时候, 懒初始化
	listDemo.PushBack(1)
	//fmt.Println(listDemo.Len())
	listDemo.Init() //清空
	//fmt.Println(listDemo.Len())
	listDemo.PushBack(2)
	listDemo.PushFront("1")
	for v := listDemo.Back(); v != nil; v = v.Prev() {
		fmt.Printf("%+v:%+v\n", reflect.TypeOf(v.Value), v.Value)
	}
	// ring ==>环形链表, 没有首尾, 任意一个node都可以看做是这个链表本身
	r := ring.New(5) // 创建长度为5的环形链表
	// 遍历链表赋值，环形链表的遍历比较特殊
	for i, at := 0, r.Next(); i < r.Len(); at, i = at.Next(), i+1 {
		at.Value = i
	}
	// 遍历链表的值
	for i, at := 0, r.Next(); i < r.Len(); at, i = at.Next(), i+1 {
		log.Printf("%v = %v", i, at.Value)
	}
	// 上面的遍历已经将r这个指针指向了值为4的这个元素
	log.Println("r:", r.Value)            // 打印 4
	log.Println("next:", r.Next().Value)  // 打印 0
	log.Println("prev", r.Prev().Value)   // 打印 3
	log.Println("move:", r.Move(2).Value) // 打印 1
	r.Do(func(i interface{}) { //i就是遍历了
		if utils.ParseInt64(i)%2 == 0 {
			fmt.Printf("偶数:%d\n", i)
		}
	})
	// Unlink(n) 从下一个next算, 删除指定个数, 返回删掉的list
	remove := r.Unlink(2)
	r.Link(remove)
	fmt.Println(r.Len())
}
