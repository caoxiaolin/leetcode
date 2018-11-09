package leetcode

/**
 * @title 填充同一层的兄弟节点 II
 *
 * 给定一个二叉树
 * struct TreeLinkNode {
 *   TreeLinkNode *left;
 *   TreeLinkNode *right;
 *   TreeLinkNode *next;
 * }
 *
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 *
 * 初始状态下，所有 next 指针都被设置为 NULL。
 *
 * 说明:
 * 你只能使用额外常数空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 *
 * 示例:
 * 给定二叉树，
 *
 *      1
 *    /  \
 *   2    3
 *  / \    \
 * 4   5    7
 *
 * 调用你的函数后，该二叉树变为：
 *
 *      1 -> NULL
 *    /  \
 *   2 -> 3 -> NULL
 *  / \    \
 * 4-> 5 -> 7 -> NULL
 *
 * @see https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 * @difficulty Medium
 */

func connect2(root *TreeLinkNode) {
	var queue []interface{}
	queue = append(queue, root)
	queue = append(queue, nil)
	for len(queue) > 1 {
		if queue[0] == nil {
			queue = append(queue, nil)
		} else {
			node := queue[0].(*TreeLinkNode)
			if queue[1] != nil {
				node.Next = queue[1].(*TreeLinkNode)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[1:]
	}
}
