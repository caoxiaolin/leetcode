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

    list := []int{}
	for res != nil {
		list = append(list, res.Val)
		res = res.Next
	}
    result := []int{13, 18, 20, 20, 25, 27, 32, 45, 47, 48, 52, 88}
    if SliceEqual(result, list) != true{
        t.Errorf("\n expect: %v\nbut got: %v", result, list)
    }
}
