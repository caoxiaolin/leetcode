package leetcode

import (
	"testing"
)

func TestSortedListToBST(t *testing.T) {
	list := []int{-10, -3, 1, 5, 9}
	listnode := InitListNode(list)

	res := sortedListToBST(listnode)

	result := LevelorderBT(res)
	expect := []int{1, -3, 9, -10, 5}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
