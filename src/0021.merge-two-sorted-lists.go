package leetcode

/**
 * @title 合并两个有序链表
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 * @see https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 * @difficulty Easy
 *
 * 基本思路：
 * 同步遍历两个链表，比较val，对于val小的，取出来加入新链表，同时当前链表节点后移，循环比较，直至一个链表结束
 * 如果有一个链表没有遍历完，直接加到新链表末尾，结束。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	res := new(ListNode) //初始化一个新链表
	tmp := res
	for p != nil && q != nil {
		if p.Val < q.Val {
			tmp.Next = p
			p = p.Next
			tmp = tmp.Next
		} else {
			tmp.Next = q
			q = q.Next
			tmp = tmp.Next
		}
	}
	if p != nil {
		tmp.Next = p
	}
	if q != nil {
		tmp.Next = q
	}
	return res.Next
}
