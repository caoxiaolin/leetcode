package leetcode

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	s := "12345#6"
	treenode := InorderBuildBT(s)
	flatten(treenode)

	LevelorderBTResult = []int{}
	LevelorderBT(treenode)
	result := LevelorderBTResult

	fmt.Printf("%v\n", result)
}
