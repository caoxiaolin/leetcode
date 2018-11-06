package leetcode

import (
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	list := make([]interface{}, 10)
	list[0] = 1
	list[1] = 2
	list[2] = 3
	list[3] = 4
	list[4] = 5
	list[5] = 6
	list[6] = nil
	list[7] = nil
	list[8] = nil
	list[9] = 7
	tree := LevelorderBuildBT(list)
	result := findBottomLeftValue(tree)
	expect := 7
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
