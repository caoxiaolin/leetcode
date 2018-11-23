package leetcode

/**
 * @title 在每个树行中找最大值
 *
 * 您需要在二叉树的每一行中找到最大的值。
 *
 * 示例：
 *
 * 输入:
 *
 *           1
 *          / \
 *         3   2
 *        / \   \
 *       5   3   9
 *
 * 输出: [1, 3, 9]
 *
 * @see https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/description/
 * @difficulty Medium
 *
 * 基本思路：
 * 层序遍历，找每一行的最大值
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var queue []interface{}
	var res []int
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0].(*TreeNode)
		max := node.Val
		size := len(queue)
		for i := 0; i < size; i++ {
			node = queue[0].(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		res = append(res, max)
	}
	return res
}
