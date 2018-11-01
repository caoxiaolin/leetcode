package leetcode

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	str := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "-"}
	fmt.Printf("tokens = %v\n", str)
	result := evalRPN(str)

	expect := 12
	if result != expect {
		t.Errorf("\n expect: %d\nbut got: %d", expect, result)
	}
}
