package utils

import (
	"testing"
)

func TestIf(t *testing.T) {
	t.Logf("%+v", If(ParseInt64(1) > ParseInt64(2), 1, 2))
}
