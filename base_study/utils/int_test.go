package utils

import (
	"math"
	"testing"
)

func TestParseInt64(t *testing.T) {
	var testData = []interface {
	}{
		111,
		"222",
		"3.3",
		3.3,
	}
	for _, v := range testData {
		result := ParseInt64(v)
		t.Logf("result: %v, testData: %v", result, v)
	}
}

func TestCompare(t *testing.T) {
	var testData = []interface {
	}{
		math.MaxFloat64 + float64(100),
		222.2,
		math.MaxInt64 + uint64(100),
		uint64(0),
		true,
		false,
	}
	t.Log(Compare(testData[0], testData[1]))
	t.Log(Compare(testData[2], testData[3]))
	t.Log(Compare(testData[4], testData[5]))
}
