package leetcode

/**
 * @title 二叉树展开为链表
 *
 * 给定一个二叉树，原地将它展开为链表。
 *
 * 例如，给定二叉树
 *
 *     1
 *    / \
 *   2   5
 *  / \   \
 * 3   4   6
 *
 * 将其展开为：
 * 1
 *  \
 *   2
 *    \
 *     3
 *      \
 *       4
 *        \
 *         5
 *          \
 *           6
 *
 * @see https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
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
func flatten(root *TreeNode) {
	p := root
	for p != nil {
		//保存Right
		n := p.Right
		if p.Left != nil {
			//遍历到最终节点
			q := p.Left
			for q.Right != nil {
				q = q.Right
			}
			q.Right = n
			p.Right = p.Left
			p.Left = nil
		}
		p = p.Right
	}
}
