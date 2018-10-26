package leetcode

import (
	"fmt"
	"testing"
)

func TestPruneTree(t *testing.T) {
	s := "1000101"
	treenode := InorderBuildBT(s)
	tree := pruneTree(treenode)

	InorderBTResult = []int{}
	InorderBT(tree)
	result := InorderBTResult

	fmt.Printf("%v\n", result)
}
