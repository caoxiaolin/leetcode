package leetcode

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
	list := make([]interface{}, 7)
	list[0] = 3
	list[1] = 9
	list[2] = 20
	list[3] = nil
	list[4] = nil
	list[5] = 15
	list[6] = 7
	tree := LevelorderBuildBT(list)

	result := levelOrder(tree)
	expect := [][]int{{3}, {9, 20}, {15, 7}}
	for i, v := range result {
		if SliceEqual(v, expect[i]) != true {
			t.Errorf("\n expect: %v\nbut got: %v", expect, result)
		}
	}

	tree = nil
	result = levelOrder(tree)
	if len(result) > 0 {
		t.Errorf("\n expect: %v\nbut got: xxx", [][]int{})
	}
}
