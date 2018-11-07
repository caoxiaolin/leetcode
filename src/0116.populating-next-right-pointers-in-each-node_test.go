package leetcode

import (
	"testing"
)

func TestConnect(t *testing.T) {
	//初始化一棵树
	tree := new(TreeLinkNode)
	tree.Left = new(TreeLinkNode)
	tree.Right = new(TreeLinkNode)
	tree.Left.Left = new(TreeLinkNode)
	tree.Left.Right = new(TreeLinkNode)
	tree.Right.Left = new(TreeLinkNode)
	tree.Right.Right = new(TreeLinkNode)
	connect(tree)
	result := tree.Left.Next
	expect := tree.Right
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
	result = tree.Left.Right.Next
	expect = tree.Right.Left
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
	if tree.Right.Right.Next != nil {
		t.Errorf("\n expect: nil\nbut got: %v", tree.Right.Right.Next)
	}
}
