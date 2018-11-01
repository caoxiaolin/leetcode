package leetcode

import (
	"testing"
)

func TestReorderList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	listnode := InitListNode(list)

	reorderList(listnode)

	result := ListNode2Slice(listnode)
	expect := []int{1, 6, 2, 5, 3, 4}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	list = []int{1}
	listnode = InitListNode(list)
	reorderList(listnode)
	result = ListNode2Slice(listnode)
	expect = []int{1}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
