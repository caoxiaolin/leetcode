package leetcode

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	n := 26
	fmt.Printf("n = %d\n", n)
	result := convertToTitle(n)

	expect := "Z"
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
