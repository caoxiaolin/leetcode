package leetcode

/**
 * @title 有序链表转换二叉搜索树
 *
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定的有序链表： [-10, -3, 0, 5, 9],
 *
 * 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
 *
 *       0
 *      / \
 *    -3   9
 *    /   /
 *  -10  5
 *
 * @see https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * 基本思路：
 * 快慢指针，二分查找中间节点作为root，左右链表递归
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	p, q, tmp := head, head, head
	//快慢指正找中间点
	for q != nil && q.Next != nil {
		tmp = p
		p = p.Next
		q = q.Next.Next
	}
	left := head
	if tmp.Next == nil { //只剩一个节点的时候
		left = nil
	} else {
		tmp.Next = nil //断开left和节点的连接
	}
	right := p.Next
	t := new(TreeNode)
	t.Val = p.Val
	if left != nil {
		t.Left = sortedListToBST(left)
	}
	if right != nil {
		t.Right = sortedListToBST(right)
	}
	return t
}
