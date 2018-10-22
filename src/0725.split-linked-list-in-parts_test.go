package leetcode

import (
	"fmt"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6}
	listnode := InitListNode(list)

	k := 3
	fmt.Printf("k = %d\n", k)
	res := splitListToParts(listnode, k)

	fmt.Println(res)
}
