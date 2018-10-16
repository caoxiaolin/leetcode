package leetcode

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

    node := listnode.Next.Next.Next.Next
    deleteNode(node)

	result := ListNode2Slice(listnode)
	expect := []int{43, 20, 15, 7, 48, 3, 67, 80, 27}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
