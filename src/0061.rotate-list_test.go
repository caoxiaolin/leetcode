package leetcode

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	k := 13
	fmt.Printf("k = %d\n", k)
	res := rotateRight(listnode, k)

	result := ListNode2Slice(res)
	expect := []int{67, 80, 27, 43, 20, 15, 7, 22, 48, 3}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	list = []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode = InitListNode(list)

	k = 0
	fmt.Printf("k = %d\n", k)
	res = rotateRight(listnode, k)

	result = ListNode2Slice(res)
	expect = []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
