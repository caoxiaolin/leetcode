package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	list := []int{5, 7, 7, 8, 8, 10}

	target := 8
	fmt.Printf("target = %d\n", target)
	result := searchRange(list, target)

	expect := []int{3, 4}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	target = 6
	fmt.Printf("target = %d\n", target)
	result = searchRange(list, target)

	expect = []int{-1, -1}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	target = 3
	fmt.Printf("target = %d\n", target)
	result = searchRange(list, target)

	expect = []int{-1, -1}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

}
