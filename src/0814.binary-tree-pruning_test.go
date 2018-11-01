package leetcode

import (
	"fmt"
	"testing"
)

func TestPruneTree(t *testing.T) {
	list := make([]interface{}, 5)
	list[0] = 1
	list[1] = nil
	list[2] = 0
	list[3] = 0
	list[4] = 1
	tree := LevelorderBuildBT(list)

	result := pruneTree(tree)

	fmt.Printf("%v\n", LevelorderBT(result))
}
