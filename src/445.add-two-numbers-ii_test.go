package leetcode

import (
	"testing"
)

func TestAddTwoNumbers2(t *testing.T) {
	list1 := []int{5, 7, 2, 8, 3}
	listnode1 := InitListNode(list1)

	list2 := []int{2, 9, 2, 1, 1, 8}
	listnode2 := InitListNode(list2)

	res := addTwoNumbers2(listnode1, listnode2)

	result := ListNode2Slice(res)
	expect := []int{3, 4, 9, 4, 0, 1}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
