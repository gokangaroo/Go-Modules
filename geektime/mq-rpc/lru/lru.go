package main

import (
	"container/list"
	"fmt"
)

/*
lru算法模拟, 使用list来记录, 新数据被访问后都移动到链表的后端, 当淘汰的时候优先淘汰,开头的数据
> 对于golang而言, container的链表, 头尾实际是一样的, 对于删除而言, 当然为了与其他语言一致,还是删除开头
而map的作用是为了快速访问, list主要是为了记录淘汰
优点: 实现简单.
> LRU的改进算法: https://www.jianshu.com/p/c4e4d55706ff
缺点: 一份数据存了两份: 我看错了, 实际map存的是地址, 那么这个没有问题
*/

//func NewLRUCache(cap int)(*LRUCache)
//func (lru *LRUCache)Set(k,v interface{})(error)
//func (lru *LRUCache)Get(k interface{})(v interface{},ret bool,err error)
//func (lru *LRUCache)Remove(k interface{})(bool)

import (
	"errors"
)

// https://studygolang.com/articles/4008

func main() {
	cache := NewLRUCache(5)
	cache.Set(1, 1)
	cache.Set(2, 2)
	cache.Set(3, 3)
	cache.Set(4, 4)
	cache.Set(5, 5)
	fmt.Println(cache.Size())
	cache.Get(1)
	cache.Set(6, 6)
	fmt.Println(cache.Size())
	for front := cache.dlist.Front(); front != nil; front = front.Next() {
		fmt.Println(front.Value)
	}
}

type CacheNode struct {
	Key, Value interface{}
}

func (cnode *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{k, v}
}

type LRUCache struct {
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element)}
}

func (lru *LRUCache) Size() int {
	return lru.dlist.Len()
}

func (lru *LRUCache) Set(k, v interface{}) error {
	if lru.dlist == nil {
		return errors.New("LRUCache结构体未初始化")
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToBack(pElement)
		pElement.Value.(*CacheNode).Value = v
		return nil
	}

	newElement := lru.dlist.PushBack(&CacheNode{k, v})
	lru.cacheMap[k] = newElement

	if lru.dlist.Len() > lru.Capacity {
		//移掉链表第一个
		lastElement := lru.dlist.Front()
		if lastElement == nil {
			return nil
		}
		cacheNode := lastElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.Key)
		lru.dlist.Remove(lastElement)
	}
	return nil
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
	if lru.cacheMap == nil {
		return v, false, errors.New("LRUCache结构体未初始化")
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToBack(pElement)
		return pElement.Value.(*CacheNode).Value, true, nil
	}
	return v, false, nil
}

func (lru *LRUCache) Remove(k interface{}) bool {
	if lru.cacheMap == nil {
		return false
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.Key)
		lru.dlist.Remove(pElement)
		return true
	}
	return false
}
