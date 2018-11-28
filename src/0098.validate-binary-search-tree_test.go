package leetcode

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	list := make([]interface{}, 7)
	list[0] = 5
	list[1] = 1
	list[2] = 4
	list[3] = nil
	list[4] = nil
	list[5] = 3
	list[6] = 6
	tree := LevelorderBuildBT(list)

	result := isValidBST(tree)
	expect := false
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
