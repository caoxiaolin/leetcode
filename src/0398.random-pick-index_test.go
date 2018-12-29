package leetcode

import (
	"fmt"
	"testing"
)

func TestPick(t *testing.T) {
	nums := []int{44, 20, 15, 43, 22, 48, 43, 67, 43, 27}
	obj := Constructor0398(nums)
	target := 43
	result := obj.Pick(target)
	fmt.Println(result)
}
