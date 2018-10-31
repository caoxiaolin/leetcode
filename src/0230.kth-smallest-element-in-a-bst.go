package leetcode

/**
 * 二叉搜索树中第K小的元素
 *
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 *
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 *
 * 示例 1:
 * 输入: root = [3,1,4,null,2], k = 1
 *    3
 *   / \
 *  1   4
 *   \
 *    2
 * 输出: 1
 *
 * 示例 2:
 * 输入: root = [5,3,6,2,4,null,null,1], k = 3
 *        5
 *       / \
 *      3   6
 *     / \
 *    2   4
 *   /
 *  1
 * 输出: 3
 *
 * 进阶：
 * 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
 *
 * @see https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 * @difficulty Medium
 *
 * 基本思路：
 * 中序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	num := 0 //遍历次数比较
	res := 0 //节点Val
	getKthSmallest(root, k, &num, &res)
	return res
}

func getKthSmallest(root *TreeNode, k int, num *int, res *int) {
	//*res>0，说明已经找到节点，直接返回
	if root == nil || *res > 0 {
		return
	}
	getKthSmallest(root.Left, k, num, res)
	*num++
	if *num == k {
		*res = root.Val
		return
	}
	getKthSmallest(root.Right, k, num, res)
}
