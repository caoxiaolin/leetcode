package leetcode

import (
	"testing"
)

func TestLargestValues(t *testing.T) {
	list := make([]interface{}, 7)
	list[0] = 1
	list[1] = 3
	list[2] = 2
	list[3] = 5
	list[4] = 3
	list[5] = nil
	list[6] = 9
	tree := LevelorderBuildBT(list)

	result := largestValues(tree)
	expect := []int{1, 3, 9}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	tree = nil
	result = largestValues(tree)
	if len(result) > 0 {
		t.Errorf("\n expect: %v\nbut got: xxx", []int{})
	}
}
