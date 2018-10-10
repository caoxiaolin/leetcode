package leetcode

import (
	"fmt"
    "testing"
)

func TestPartition(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
    listnode := InitListNode(list)

	x := 30
	res := partition(listnode, x)

	fmt.Printf("x = %d\n", x)
    list1 := []int{}
	for res != nil {
		list1 = append(list1, res.Val)
		res = res.Next
	}
    result := []int{20, 15, 7, 22, 3, 27, 43, 48, 67, 80}
    if SliceEqual(result, list1) != true{
        t.Errorf("\n expect: %v\nbut got: %v", result, list1)
    }
}
