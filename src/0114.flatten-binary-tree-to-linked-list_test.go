package leetcode

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	s := "12345#6"
	treenode := PreorderBuildBT(s)
	flatten(treenode)

	result := LevelorderBT(treenode)

	fmt.Printf("%v\n", result)
}
