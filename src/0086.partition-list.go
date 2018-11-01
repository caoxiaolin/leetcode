package leetcode

/**
 * @title 分隔链表
 *
 * 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
 *
 * 你应当保留两个分区中每个节点的初始相对位置。
 *
 * 示例:
 * 输入: head = 1->4->3->2->5->2, x = 3
 * 输出: 1->2->2->4->3->5
 *
 * @see https://leetcode-cn.com/problems/partition-list/description/
 * @difficulty Medium
 *
 * 基本思路：
 * 建两个链表 left 和 right，遍历给定的链表，节点与 x 比较，小于 x 的节点加到left，否则加到right，最后合并
 */

func partition(head *ListNode, x int) *ListNode {
	left := &ListNode{Val: 0, Next: nil}
	right := &ListNode{Val: 0, Next: nil}
	lefthead, righthead := left, right
	p := head
	for p != nil {
		//小于 x ，增加到left
		if p.Val < x {
			left.Next = &ListNode{Val: p.Val, Next: nil}
			left = left.Next
		} else {
			right.Next = &ListNode{Val: p.Val, Next: nil}
			right = right.Next
		}
		p = p.Next
	}
	left.Next = righthead.Next //跳过构造的right的头结点
	return lefthead.Next
}
