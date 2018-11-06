package leetcode

import "strconv"

var StrBT string //构造二叉树的字符串

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
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/**
 * merge two slice
 */
/*
func mergeSlice(a, b []int) []int {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	for _, v := range b {
		a = append(a, v)
	}
	return a
}
*/

/**
 * 前序构建二叉树
 */
func PreorderBuildBT(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	c := string(s[0])
	StrBT = s[1:]
	if c != "#" {
		var t TreeNode
		val, _ := strconv.Atoi(c)
		t.Val = val
		t.Left = PreorderBuildBT(StrBT)
		t.Right = PreorderBuildBT(StrBT)
		return &t
	}
	return nil
}

/**
 * 层序构建二叉树
 */
func LevelorderBuildBT(a []interface{}) *TreeNode {
	length := len(a)
	if length == 0 {
		return nil
	}
	queue := make([]interface{}, length)
	i, j := 0, 0
	root := &TreeNode{a[0].(int), nil, nil}
	queue[0] = root
	for i < length {
		if queue[j] == nil {
		}
		if queue[j] != nil {
			node := queue[j].(*TreeNode)
			i++
			if i == length {
				break
			}
			if a[i] != nil {
				node.Left = &TreeNode{a[i].(int), nil, nil}
				queue[i] = node.Left
			} else {
				queue[i] = nil
			}
			i++
			if i == length {
				break
			}
			if a[i] != nil {
				node.Right = &TreeNode{a[i].(int), nil, nil}
				queue[i] = node.Right
			} else {
				queue[i] = nil
			}
		}
		j++
	}
	return root
}

/**
 * 层序遍历二叉树
 */
func LevelorderBT(root *TreeNode) []int {
	queue := make([]*TreeNode, 20)
	i, j := 0, 1
	queue[i] = root
	res := []int{}
	for queue[i] != nil {
		node := queue[i]
		res = append(res, node.Val)
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
	return res
}

/**
 * 前序遍历二叉树
 */
/*
func PreorderBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
	}
	res = mergeSlice(res, PreorderBT(root.Left))
	res = mergeSlice(res, PreorderBT(root.Right))
	return res
}
*/
