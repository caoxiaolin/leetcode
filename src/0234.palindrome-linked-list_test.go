package leetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	list := []int{1, 2, 3, 2, 1}
	listnode := InitListNode(list)
	res := isPalindrome(listnode)
	if res != true {
		t.Errorf("\n expect: true\nbut got: %v", res)
	}

	list = []int{1, 2}
	listnode = InitListNode(list)
	res = isPalindrome(listnode)
	if res != false {
		t.Errorf("\n expect: false\nbut got: %v", res)
	}

	list = []int{1}
	listnode = InitListNode(list)
	res = isPalindrome(listnode)
	if res != true {
		t.Errorf("\n expect: false\nbut got: %v", res)
	}
}
