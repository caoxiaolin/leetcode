package leetcode

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	list := []int{2, 0, 2, 1, 1, 0}

	sortColors(list)

	expect := []int{0, 0, 1, 1, 2, 2}
	if SliceEqual(list, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, list)
	}
}
