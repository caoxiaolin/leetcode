package leetcode

import (
	"fmt"
	"testing"
)

func TestNumComponents(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}
	listnode := InitListNode(list)

	G := []int{20, 15, 48, 80}
	fmt.Printf("G = %v\n", G)
	result := numComponents(listnode, G)

	expect := 3
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
