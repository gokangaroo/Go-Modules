package main

import (
	"github.com/willf/bitset"
)

const DefaultSize = 2 << 24

var seeds = []uint{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	set       *bitset.BitSet
	functions [6]SimpleHash
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(bf.functions); i++ {
		bf.functions[i] = SimpleHash{DefaultSize, seeds[i]}
	}
	bf.set = bitset.New(DefaultSize)
	return bf
}

//添加元素
func (bf *BloomFilter) add(value string) {
	for _, f := range bf.functions {
		bf.set.Set(f.hash(value))
	}
}

//是否包含元素
func (bf *BloomFilter) contains(value string) bool {
	if value == "" {
		return false
	}
	ret := true
	for _, f := range bf.functions {
		ret = ret && bf.set.Test(f.hash(value))
	}
	return ret
}

type SimpleHash struct {
	cap  uint
	seed uint
}

func (s *SimpleHash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.seed + uint(value[i])
	}
	return (s.cap - 1) & result
}
