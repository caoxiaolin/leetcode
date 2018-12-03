package leetcode

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	result := canCompleteCircuit(gas, cost)

	expect := 3
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
