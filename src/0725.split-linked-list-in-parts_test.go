package leetcode

import (
	"fmt"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6}
	listnode := InitListNode(list)

	k := 1
	fmt.Printf("k = %d\n", k)
	res := splitListToParts(listnode, k)
	fmt.Println(res)

	k = 2
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

	k = 3
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

	k = 4
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

	k = 5
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

	k = 6
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

	k = 7
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

	k = 10
	fmt.Printf("k = %d\n", k)
	res = splitListToParts(listnode, k)
	fmt.Println(res)

}
