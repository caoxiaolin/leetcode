package leetcode

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	list := make([]interface{}, 5)
	list[0] = 4
	list[1] = 9
	list[2] = 0
	list[3] = 5
	list[4] = 1
	tree := LevelorderBuildBT(list)

	result := sumNumbers(tree)
	expect := 1026
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
