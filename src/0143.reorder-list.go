package leetcode

/**
 * @title 重排链表
 *
 * 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
 * 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * 示例 1:
 *
 * 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 * 示例 2:
 *
 * 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
 *
 * @see https://leetcode-cn.com/problems/reorder-list/description/
 *
 * 基本思路：
 * 找到链表中点，一分为二，后半段反转，同时遍历两个链表，隔一个插一个节点，形成新的链表
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}

	//分段，left长度>=right长度
	right := p.Next
	p.Next = nil
	left := head

	//反转
	first := 1
	tmp := right
	for right != nil {
		rn := right.Next
		if first == 1 {
			right.Next = nil
			first = 0
		} else {
			right.Next = tmp
		}
		tmp = right
		right = rn
	}
	right = tmp

	//遍历插入
	for right != nil {
		lnext := left.Next
		rnext := right.Next
		left.Next = right
		right.Next = lnext
		left = lnext
		right = rnext
	}
}
