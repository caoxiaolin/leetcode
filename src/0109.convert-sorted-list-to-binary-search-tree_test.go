package leetcode

import (
	"testing"
)

func TestSortedListToBST(t *testing.T) {
	list := []int{-10, -3, 1, 5, 9}
	listnode := InitListNode(list)

	res := sortedListToBST(listnode)

	InorderBTResult = []int{}
	InorderBT(res)

	result := InorderBTResult
	expect := []int{-10, -3, 1, 5, 9}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
