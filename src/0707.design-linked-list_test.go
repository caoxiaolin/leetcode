package leetcode

import (
	"testing"
)

func TestDesignLinkedList(t *testing.T) {
	obj := Constructor707()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	result1 := obj.Get(1)
	obj.DeleteAtIndex(1)
	result2 := obj.Get(1)

	expect1, expect2 := 2, 3
	if result1 != expect1 {
		t.Errorf("\n expect: %v\nbut got: %v", expect1, result1)
	}
	if result2 != expect2 {
		t.Errorf("\n expect: %v\nbut got: %v", expect2, result2)
	}
}
