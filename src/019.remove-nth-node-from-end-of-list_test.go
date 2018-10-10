package leetcode

import (
	"fmt"
    "testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
    listnode := InitListNode(list)

	n := 4
	res := removeNthFromEnd(listnode, n)

	fmt.Printf("n = %d\n", n)
    list1 := []int{}
	for res != nil {
		list1 = append(list1, res.Val)
		res = res.Next
	}
    result := []int{43, 20, 15, 7, 22, 48, 67, 80, 27}
    if SliceEqual(result, list1) != true{
        t.Errorf("\n expect: %v\nbut got: %v", result, list1)
    }
}
