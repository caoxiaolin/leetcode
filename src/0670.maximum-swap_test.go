package leetcode

import (
	"fmt"
	"testing"
)

func TestMaximumSwap(t *testing.T) {
	num := 1993
	fmt.Printf("num = %d\n", num)
	result := maximumSwap(num)

	expect := 9913
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
