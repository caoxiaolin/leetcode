package leetcode

import (
	"testing"
)

func TestSortList(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	res := sortList(listnode)

	result := ListNode2Slice(res)
	expect := []int{3, 7, 15, 20, 22, 27, 43, 48, 67, 80}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
