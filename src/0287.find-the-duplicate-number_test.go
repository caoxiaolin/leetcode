package leetcode

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	list := []int{3, 1, 3, 4, 2}

	result := findDuplicate(list)

	expect := 3
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
