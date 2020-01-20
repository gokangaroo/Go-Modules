package utils

import (
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	var testData = []interface {
	}{
		111,
		3.3,
		-111,
		"233 ",
		true,
		[]string{"3.3"},
		// func (d Weekday) String() string {...}
		time.Now().Weekday(),
	}
	for _, v := range testData {
		result := ToString(v)
		t.Logf("testData: %+v, result: %s", v, result)
	}
}
