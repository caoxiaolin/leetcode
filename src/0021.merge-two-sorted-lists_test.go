package leetcode

import (
	"testing"
)

func TestMergeTowLists(t *testing.T) {
	list1 := []int{18, 20, 25, 27, 32, 48}
	listnode1 := InitListNode(list1)

	list2 := []int{13, 20, 45, 47, 52, 88}
	listnode2 := InitListNode(list2)

	res := mergeTwoLists(listnode1, listnode2)

	result := ListNode2Slice(res)
	expect := []int{13, 18, 20, 20, 25, 27, 32, 45, 47, 48, 52, 88}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	list1 = []int{18, 20, 25, 27, 32, 48, 92, 100}
	listnode1 = InitListNode(list1)

	list2 = []int{13, 20, 45, 47, 52, 88}
	listnode2 = InitListNode(list2)

	res = mergeTwoLists(listnode1, listnode2)

	result = ListNode2Slice(res)
	expect = []int{13, 18, 20, 20, 25, 27, 32, 45, 47, 48, 52, 88, 92, 100}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

}
