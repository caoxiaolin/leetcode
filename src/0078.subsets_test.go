package leetcode

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	list := []int{9, 0, 3, 5, 7}
	result := subsets(list)

	fmt.Println(result)
}
