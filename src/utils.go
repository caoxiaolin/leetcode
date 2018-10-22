package leetcode

import "strconv"

var StrBT string                 //构造二叉树的字符串
var InorderBTResult = []int{}    //中序遍历二叉树结果
var LevelorderBTResult = []int{} //层序遍历二叉树结果

/**
 * Init singly-linked list
 */
func InitListNode(list []int) *ListNode {
	nodehead := &ListNode{Val: 0, Next: nil}
	listnode := nodehead
	for _, v := range list {
		listnode.Next = &ListNode{Val: v, Next: nil}
		listnode = listnode.Next
	}
	return nodehead.Next
}

/**
 * ListNode to Slice
 */
func ListNode2Slice(list *ListNode) []int {
	res := []int{}
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return res
}

/**
 * Equal two slice
 */
func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/**
 * 前序构建二叉树
 */
func InorderBuildBT(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	c := string(s[0])
	StrBT = s[1:]
	if c != "#" {
		var t TreeNode
		val, _ := strconv.Atoi(c)
		t.Val = val
		t.Left = InorderBuildBT(StrBT)
		t.Right = InorderBuildBT(StrBT)
		return &t
	}
	return nil
}

/**
 * 层序遍历二叉树
 */
func LevelorderBT(root *TreeNode) {
	queue := make([]*TreeNode, 20)
	i, j := 0, 1
	queue[i] = root
	for queue[i] != nil {
		node := queue[i]
		LevelorderBTResult = append(LevelorderBTResult, node.Val)
		i++
		if node.Left != nil {
			queue[j] = node.Left
			j++
		}
		if node.Right != nil {
			queue[j] = node.Right
			j++
		}
	}
}

/**
 * 前序遍历二叉树
 */
func InorderBT(root *TreeNode) {
	if root == nil {
		return
	}
	if root != nil {
		InorderBTResult = append(InorderBTResult, root.Val)
	}
	InorderBT(root.Left)
	InorderBT(root.Right)
}
