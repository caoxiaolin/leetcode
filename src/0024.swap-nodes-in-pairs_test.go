package leetcode

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	res := swapPairs(listnode)

	result := ListNode2Slice(res)
	expect := []int{20, 43, 7, 15, 48, 22, 67, 3, 27, 80}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	list = []int{1}
	listnode = InitListNode(list)

	res = swapPairs(listnode)

	result = ListNode2Slice(res)
	expect = []int{1}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
