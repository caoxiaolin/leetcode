package leetcode

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	path := "/a/./b/../../c/"
	result := simplifyPath(path)
	expect := "/c"
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
