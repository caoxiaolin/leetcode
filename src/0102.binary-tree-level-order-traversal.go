package leetcode

/**
 * @title 二叉树的层次遍历
 *
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回其层次遍历结果：
 * [
 *   [3],
 *   [9,20],
 *   [15,7]
 * ]
 *
 * @see https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []interface{}
	var result [][]int
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		res := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i].(*TreeNode)
			res[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, res)
		queue = queue[size:]
	}
	return result
}
