package leetcode

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	x := 30
	fmt.Printf("x = %d\n", x)
	res := partition(listnode, x)

	result := ListNode2Slice(res)
	expect := []int{20, 15, 7, 22, 3, 27, 43, 48, 67, 80}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
