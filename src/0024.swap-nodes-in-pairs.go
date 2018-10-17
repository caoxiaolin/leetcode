package leetcode

/**
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 * 说明:
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * @see https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &ListNode{Val: 0, Next: head}
	phead := tmp
	for tmp.Next != nil && tmp.Next.Next != nil {
		p := tmp.Next
		q := tmp.Next.Next
		//swap
		p.Next = q.Next
		q.Next = p
		tmp.Next = q

		//后移
		tmp = p
	}
	return phead.Next
}