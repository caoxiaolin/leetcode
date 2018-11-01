package leetcode

import (
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	res := insertionSortList(listnode)

	result := ListNode2Slice(res)
	expect := []int{3, 7, 15, 20, 22, 27, 43, 48, 67, 80}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	list = []int{43}
	listnode = InitListNode(list)

	res = insertionSortList(listnode)

	result = ListNode2Slice(res)
	expect = []int{43}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

}
