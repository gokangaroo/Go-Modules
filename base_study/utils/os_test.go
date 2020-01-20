package utils

import (
	"testing"
)

func TestCheckFileIfExist(t *testing.T) {
	var testData = []string{
		`/home/huijia`,
	}
	for _, v := range testData {
		result := CheckFileIfExist(v)
		t.Logf("testData: %s, result: %v", v, result)
	}
}
