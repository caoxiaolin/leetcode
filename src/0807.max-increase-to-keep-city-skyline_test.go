package leetcode

import (
	"testing"
)

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	result := maxIncreaseKeepingSkyline(grid)

	expect := 35
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
