package leetcode

import (
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	result := validateStackSequences(pushed, popped)
	expect := false
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
