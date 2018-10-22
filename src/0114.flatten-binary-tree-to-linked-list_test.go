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
	/*
	   result := LevelorderBTResult
	   expect := []int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}
	   if SliceEqual(result, expect) != true {
	       t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	   }
	*/
}
