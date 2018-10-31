package leetcode

/**
 * @title 二叉树的中序遍历
 *
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 *    1
 *     \
 *      2
 *     /
 *    3
 *
 * 输出: [1,3,2]
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 * @see https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 * @difficulty Medium
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = func(a, b []int) []int {
		for _, v := range b {
			a = append(a, v)
		}
		return a
	}(res, inorderTraversal(root.Left))
	if root != nil {
		res = append(res, root.Val)
	}
	res = func(a, b []int) []int {
		for _, v := range b {
			a = append(a, v)
		}
		return a
	}(res, inorderTraversal(root.Right))
	return res
}
