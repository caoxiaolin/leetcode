package leetcode

import (
	"fmt"
    "testing"
)

func TestReverseBetween(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
    listnode := InitListNode(list)

	m, n := 4, 7
	fmt.Printf("m = %d\tn = %d\n", m, n)
	res := reverseBetween(listnode, m, n)

    result := ListNode2Slice(res)
    expect := []int{43, 20, 15, 3, 48, 22, 7, 67, 80, 27}
    if SliceEqual(result, expect) != true{
        t.Errorf("\n expect: %v\nbut got: %v", expect, result)
    }
}
