package leetcode

import (
	"testing"
)

func TestMinAddToMakeValid(t *testing.T) {
	S := "()))(("
	result := minAddToMakeValid(S)

	expect := 4
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
