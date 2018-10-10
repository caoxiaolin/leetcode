package leetcode

import (
	"fmt"
    "testing"
)

func TestRemoveElements(t *testing.T) {
	list := []int{43, 20, 15, 7, 20, 20, 20, 67, 80, 27}
    listnode := InitListNode(list)

	val := 20
	fmt.Printf("val = %d\n", val)
	res := removeElements(listnode, val)

    result := ListNode2Slice(res)
    expect := []int{43, 15, 7, 67, 80, 27}
    if SliceEqual(result, expect) != true{
        t.Errorf("\n expect: %v\nbut got: %v", expect, result)
    }
}
