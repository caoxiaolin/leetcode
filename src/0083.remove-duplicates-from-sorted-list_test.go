package leetcode

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	list := []int{13, 20, 25, 27, 27, 28, 28, 28, 80, 97, 97}
	listnode := InitListNode(list)

	res := deleteDuplicates(listnode)

	result := ListNode2Slice(res)
	expect := []int{13, 20, 25, 27, 28, 80, 97}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	list = []int{13}
	listnode = InitListNode(list)

	res = deleteDuplicates(listnode)

	result = ListNode2Slice(res)
	expect = []int{13}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

}
