package leetcode

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	list := []int{43, 20, 15, 7, 22, 48, 3, 67, 80, 27}

	k := 2
	fmt.Printf("k = %d\n", k)
	result := findKthLargest(list, k)
	expect := 67
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
