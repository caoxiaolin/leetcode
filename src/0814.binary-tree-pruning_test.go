package leetcode

import (
	"fmt"
	"testing"
)

func TestPruneTree(t *testing.T) {
	s := "1000101"
	treenode := PreorderBuildBT(s)
	tree := pruneTree(treenode)

	result := PreorderBT(tree)

	fmt.Printf("%v\n", result)
}
