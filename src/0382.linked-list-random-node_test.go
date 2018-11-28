package leetcode

import (
	"fmt"
	"testing"
)

func TestGetRandom(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	obj := Constructor382(listnode)
	result := obj.GetRandom()

	fmt.Println(result)
}
