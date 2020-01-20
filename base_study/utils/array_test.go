package utils

import "testing"

func TestArrayContains(t *testing.T) {
	var testData = []struct {
		array [3]interface{}
		slice []interface{}
		val   interface{}
	}{
		{array: [...]interface{}{"0", "1", "2"}, slice: []interface{}{"1", "2"}, val: "2"},
	}
	for _, v := range testData {
		result1 := ArrayContains(v.array, v.val)
		result2 := ArrayContains(v.slice, v.val)
		t.Logf("result1: %v,result2: %v, testData: %v", result1, result2, v)
	}
}

func TestArrayCheckShortest(t *testing.T) {
	var testData = []struct {
		array [3]interface{}
		slice []interface{}
		limit int
	}{
		{array: [...]interface{}{2333, 1.2, 2233}, slice: []interface{}{2333, 1.223}, limit: 4},
	}
	for _, v := range testData {
		result1 := ArrayCheckShortest(v.array, v.limit)
		result2 := ArrayCheckShortest(v.slice, v.limit)
		t.Logf("result1: %v,result2: %v, testData: %v", result1, result2, v)
	}
}

func TestArrayCheckLongest(t *testing.T) {
	var testData = []struct {
		array [3]interface{}
		slice []interface{}
		limit int
	}{
		{array: [...]interface{}{2333, 1.2, 2233}, slice: []interface{}{23333, 1.22}, limit: 4},
	}
	for _, v := range testData {
		result1 := ArrayCheckLongest(v.array, v.limit)
		result2 := ArrayCheckLongest(v.slice, v.limit)
		t.Logf("result1: %v,result2: %v, testData: %v", result1, result2, v)
	}
}
