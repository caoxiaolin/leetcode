package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	list1 := []int{2, 4, 3}
	listnode1 := InitListNode(list1)

	list2 := []int{5, 6, 4}
	listnode2 := InitListNode(list2)

	res := addTwoNumbers(listnode1, listnode2)

	result := ListNode2Slice(res)
	expect := []int{7, 0, 8}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
