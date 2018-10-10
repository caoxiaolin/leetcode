package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	n := 4
	fmt.Printf("n = %d\n", n)
	res := removeNthFromEnd(listnode, n)

	result := ListNode2Slice(res)
	expect := []int{43, 20, 15, 7, 22, 48, 67, 80, 27}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
