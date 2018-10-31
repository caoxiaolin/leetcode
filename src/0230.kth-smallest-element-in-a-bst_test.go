package leetcode

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	list := make([]interface{}, 8)
	list[0] = 5
	list[1] = 3
	list[2] = 6
	list[3] = 2
	list[4] = 4
	list[5] = nil
	list[6] = nil
	list[7] = 1
	tree := LevelorderBuildBT(list)

	k := 3
	fmt.Printf("k = %d\n", k)
	result := kthSmallest(tree, k)

	expect := 3
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
