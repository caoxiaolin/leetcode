package leetcode

/**
 * @title 找树左下角的值
 *
 * 给定一个二叉树，在树的最后一行找到最左边的值。
 *
 * 示例 1:
 *
 * 输入:
 *
 *     2
 *    / \
 *   1   3
 *
 * 输出:
 * 1
 *
 * 示例 2:
 *
 * 输入:
 *
 *         1
 *        / \
 *       2   3
 *      /   / \
 *     4   5   6
 *        /
 *       7
 *
 * 输出:
 * 7
 *
 * @see https://leetcode-cn.com/problems/find-bottom-left-tree-value/description/
 * @difficulty Medium
 *
 * 基本思路：
 * 层序遍历，从右到左
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	var queue []interface{}
	queue = append(queue, root)
	res := 0
	for len(queue) > 0 {
		node := queue[0].(*TreeNode)
		res = node.Val
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return res
}
