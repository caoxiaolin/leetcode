package leetcode

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	list := make([]interface{}, 11)
	list[0] = 3
	list[1] = 5
	list[2] = 1
	list[3] = 6
	list[4] = 2
	list[5] = 0
	list[6] = 8
	list[7] = nil
	list[8] = nil
	list[9] = 7
	list[10] = 4
	tree := LevelorderBuildBT(list)

	result := lowestCommonAncestor(tree, tree.Left, tree.Right)
	fmt.Printf("%v\n", result)

	result = lowestCommonAncestor(tree, tree.Left, tree.Left.Right.Right)
	fmt.Printf("%v\n", result)
}
