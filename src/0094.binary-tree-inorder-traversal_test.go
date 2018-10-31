package leetcode

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	list := make([]interface{}, 4)
	list[0] = 1
	list[1] = nil
	list[2] = 2
	list[3] = 3
	tree := LevelorderBuildBT(list)

	result := inorderTraversal(tree)
	expect := []int{1, 3, 2}
	if SliceEqual(result, expect) != true {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
