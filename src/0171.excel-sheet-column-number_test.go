package leetcode

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	s := "Z"
	fmt.Printf("s = %s\n", s)
	result := titleToNumber(s)

	expect := 26
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
